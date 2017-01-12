package spotify

import (
	"encoding/json"
	"errors"
	"log"

	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"github.com/zmb3/spotify"
)

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

// GetAudioFeatures searches spotify for the user playlist
func GetAudioFeatures(client *spotify.Client, ids ...spotify.ID) ([]*spotify.AudioFeatures, error) {
	features, err := client.GetAudioFeatures(ids...)
	if err != nil {
		log.Printf("GetAudioFeatures Error: %+v", err)
		var nilSlice []*spotify.AudioFeatures
		return nilSlice, err
	}
	return features, nil
}

// GetArtistGenres searches spotify for the user playlist
func GetArtistGenres(client *spotify.Client, ids ...spotify.ID) ([]*spotify.FullArtist, error) {
	artists, err := client.GetArtists(ids...)
	if err != nil {
		log.Printf("GetArtistGenres Error: %+v", err)
		var nilSlice []*spotify.FullArtist
		return nilSlice, err
	}
	return artists, nil
}

// TracksFromPlaylist gets tracks from spotify and processes them
func TracksFromPlaylist(client *spotify.Client, playlistID spotify.ID, ownerID string) ([]trackR.Track, error) {
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
			newTrack := trackR.Track{SpotifyTrack: track.Track, SpotifyID: track.Track.SpotifyID, Features: *featureResult[0], Genres: []string{}, CustomGenres: []string{}, Playlists: []string{string(playlistID)}}
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
			newTrack := trackR.Track{SpotifyTrack: track.Track, SpotifyID: track.Track.SpotifyID}
			tracks = append(tracks, newTrack)
		}
	}
	return tracks, nil
}