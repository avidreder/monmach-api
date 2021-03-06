package spotify

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	usermw "github.com/avidreder/monmach-api/middleware/user"
	authR "github.com/avidreder/monmach-api/resources/auth"
	configR "github.com/avidreder/monmach-api/resources/config"
	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	trackR "github.com/avidreder/monmach-api/resources/track"
	"github.com/labstack/echo"

	"github.com/zmb3/spotify"
	"gopkg.in/mgo.v2/bson"
)

// LoadAuthenticator places initialized spotify client
func LoadAuthenticator(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config, err := configR.GetConfig()
		if err != nil {
			log.Printf("Could not get service config: %s", err)
		}
		file, err := os.Open(config.SpotifyCredentialsPath) // For read access.
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		credentials := authR.SpotifyCredentials{}
		err = json.Unmarshal(contents, &credentials)
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		scopes := []string{"user-read-email", "playlist-read-private", "playlist-modify-public", "playlist-modify-private"}
		auth := spotify.NewAuthenticator(config.SpotifyCallback, scopes...)
		auth.SetAuthInfo(credentials.ClientKey, credentials.Secret)
		c.Set("spotifyAuthenticator", &auth)
		log.Printf("initialize: %+v", auth.AuthURL("state"))
		return h(c)
	}
}

// GetAuthenticator retieves authenticator from the context
func GetAuthenticator(c echo.Context) *spotify.Authenticator {
	return c.Get("spotifyAuthenticator").(*spotify.Authenticator)
}

// LoadClient places initialized spotify client
func LoadClient(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config, err := configR.GetConfig()
		if err != nil {
			log.Printf("Could not get service config: %s", err)
		}
		file, err := os.Open(config.SpotifyCredentialsPath) // For read access.
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		credentials := authR.SpotifyCredentials{}
		err = json.Unmarshal(contents, &credentials)
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		auth := spotify.NewAuthenticator(config.SpotifyCallback, spotify.ScopeUserReadPrivate)
		auth.SetAuthInfo(credentials.ClientKey, credentials.Secret)

		user := usermw.GetUser(c)
		client := auth.NewClient(&user.Token)
		c.Set("spotifyClient", &client)
		return h(c)
	}
}

// GetClient retieves provider from the context
func GetClient(c echo.Context) *spotify.Client {
	return c.Get("spotifyClient").(*spotify.Client)
}

// FindDiscoverPlaylist searches spotify for the user playlist
func FindDiscoverPlaylist(client *spotify.Client) (spotify.ID, error) {
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return "", err
	}
	playlistArray := playlists.Playlists
	for _, pl := range playlistArray {
		if pl.Name == "Discover Weekly" {
			return pl.ID, nil
		}
	}
	return "", errors.New("Could not find discover playlist")
}

// FindPlaylistOwner searches spotify for the user playlist
func FindPlaylistOwner(client *spotify.Client, playlistID spotify.ID) (string, error) {
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return "", err
	}
	playlistArray := playlists.Playlists
	for _, pl := range playlistArray {
		if pl.ID == playlistID {
			return pl.Owner.ID, nil
		}
	}
	return "", errors.New("Could not find discover playlist")
}

// GetAudioFeatures searches spotify for the user playlist
func GetAudioFeatures(client *spotify.Client, ids ...spotify.ID) ([]*spotify.AudioFeatures, error) {
	var features []*spotify.AudioFeatures
	for _, id := range ids {
		feature, err := client.GetAudioFeatures(id)
		if err != nil {
			log.Printf("GetAudioFeatures Error: %+v", err)
			var nilSlice []*spotify.AudioFeatures
			return nilSlice, err
		}
		log.Print("Got feature sucessfully")
		features = append(features, feature[0])
		time.Sleep(100 * time.Millisecond)
	}
	return features, nil
}

// GetArtistGenres searches spotify for the user playlist
func GetArtistGenres(client *spotify.Client, ids ...spotify.ID) ([]*spotify.FullArtist, error) {
	var artists []*spotify.FullArtist
	for _, id := range ids {
		artist, err := client.GetArtists(id)
		if err != nil {
			log.Printf("GetArtistGenres Error: %+v", err)
			var nilSlice []*spotify.FullArtist
			return nilSlice, err
		}
		log.Print("Got artist genres sucessfully")
		artists = append(artists, artist[0])
		time.Sleep(100 * time.Millisecond)
	}
	return artists, nil
}

// TracksFromPlaylist gets tracks from spotify and processes them
func TracksFromPlaylist(client *spotify.Client, playlistID spotify.ID, ownerID string, userID bson.ObjectId) ([]trackR.Track, error) {
	tracks := []trackR.Track{}
	responseObject := []spotifyR.SpotifyResponse{}
	response, err := client.GetPlaylistTracksOpt(ownerID, playlistID, nil, "items(track(album(images(url,height,width)),name,id,artists(name,id)))")
	if err != nil {
		return []trackR.Track{}, err
	}
	responseJSON, err := json.Marshal(response.Tracks)
	if err != nil {
		return []trackR.Track{}, err
	}
	err = json.Unmarshal(responseJSON, &responseObject)
	if err != nil {
		return []trackR.Track{}, err
	}
	for _, track := range responseObject {
		featureResult, err := GetAudioFeatures(client, spotify.ID(track.Track.SpotifyID))
		if err == nil {
			newTrack := trackR.Track{SpotifyTrack: track.Track, SpotifyID: track.Track.SpotifyID, Features: *featureResult[0], Genres: make([]string, 0), CustomGenres: make([]string, 0), Playlists: []string{string(playlistID)}}
			genreSlice := []string{}
			for k, artist := range newTrack.SpotifyTrack.Artists {
				artistInfo, err := GetArtistGenres(client, spotify.ID(artist.SpotifyID))
				if err == nil {
					fullArtist := *artistInfo[0]
					newTrack.SpotifyTrack.Artists[k].Genres = fullArtist.Genres
					genreSlice = append(genreSlice, fullArtist.Genres...)
				}
				genreMap := map[string]struct{}{}
				dedupedSlice := []string{}
				for _, v := range genreSlice {
					_, ok := genreMap[v]
					if !ok {
						dedupedSlice = append(dedupedSlice, v)
						genreMap[v] = struct{}{}
					}
				}
				newTrack.Genres = dedupedSlice
			}
			newTrack.OwnerID = userID
			tracks = append(tracks, newTrack)
		} else {
			newTrack := trackR.Track{OwnerID: userID, SpotifyTrack: track.Track, SpotifyID: track.Track.SpotifyID, Genres: make([]string, 0), CustomGenres: make([]string, 0), Playlists: []string{string(playlistID)}}
			tracks = append(tracks, newTrack)
		}
	}
	return tracks, nil
}

// RecommendedTracks gets tracks from spotify and processes them
func RecommendedTracks(client *spotify.Client, params RecommendedTrackParams, userID bson.ObjectId) ([]trackR.Track, error) {
	tracks := []trackR.Track{}
	seedCount := 0
	seeds := spotify.Seeds{}
	for _, v := range params.Artists {
		if seedCount < 5 {
			seeds.Artists = append(seeds.Artists, spotify.ID(v))
			seedCount++
		}
	}
	for _, v := range params.Tracks {
		if seedCount < 5 {
			seeds.Tracks = append(seeds.Tracks, spotify.ID(v))
			seedCount++
		}
	}
	for _, v := range params.Genres {
		if seedCount < 5 {
			seeds.Genres = append(seeds.Genres, v)
			seedCount++
		}
	}
	responseObject := spotify.Recommendations{}
	response, err := client.GetRecommendations(seeds, nil, nil)
	if err != nil {
		return []trackR.Track{}, err
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return []trackR.Track{}, err
	}
	err = json.Unmarshal(responseJSON, &responseObject)
	if err != nil {
		return []trackR.Track{}, err
	}
	for _, track := range responseObject.Tracks {
		trackContainer := spotifyR.SpotifyTrack{}
		trackJSON, err := json.Marshal(track)
		if err != nil {
			log.Printf("track marshal error %+v", err)
			return []trackR.Track{}, err
		}
		err = json.Unmarshal(trackJSON, &trackContainer)
		if err != nil {
			log.Printf("track unmarshal error %+v", err)
			return []trackR.Track{}, err
		}
		log.Printf("got track with id: %+v", trackContainer.SpotifyID)
		featureResult, err := GetAudioFeatures(client, spotify.ID(trackContainer.SpotifyID))
		if err == nil {
			newTrack := trackR.Track{OwnerID: userID, SpotifyTrack: trackContainer, SpotifyID: trackContainer.SpotifyID, Features: *featureResult[0], Genres: make([]string, 0), CustomGenres: make([]string, 0), Playlists: make([]string, 0)}
			genreSlice := []string{}
			for k, artist := range newTrack.SpotifyTrack.Artists {
				artistInfo, err := GetArtistGenres(client, spotify.ID(artist.SpotifyID))
				if err == nil {
					fullArtist := *artistInfo[0]
					newTrack.SpotifyTrack.Artists[k].Genres = fullArtist.Genres
					genreSlice = append(genreSlice, fullArtist.Genres...)
				}
				genreMap := map[string]struct{}{}
				dedupedSlice := []string{}
				for _, v := range genreSlice {
					_, ok := genreMap[v]
					if !ok {
						dedupedSlice = append(dedupedSlice, v)
						genreMap[v] = struct{}{}
					}
				}
				newTrack.Genres = dedupedSlice
			}
			tracks = append(tracks, newTrack)
		} else {
			newTrack := trackR.Track{OwnerID: userID, SpotifyTrack: trackContainer, SpotifyID: trackContainer.SpotifyID, Genres: make([]string, 0), CustomGenres: make([]string, 0), Playlists: make([]string, 0)}
			tracks = append(tracks, newTrack)
		}
	}
	return tracks, nil
}

type RecommendedTrackParams struct {
	Artists []string `json:"artists"`
	Tracks  []string `json:"tracks"`
	Genres  []string `json:"genres"`
}
