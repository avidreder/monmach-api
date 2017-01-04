package spotify

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	queuemw "github.com/avidreder/monmach-api/middleware/queue"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"github.com/labstack/echo"
	"github.com/zmb3/spotify"
)

// UserPlaylists gets a user's spotify playlists
func UserPlaylists(c echo.Context) error {
	client := spotifyR.GetClient(c)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, playlists)
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

// DiscoverPlaylist gets and stores a user's spotify discover playlist
func DiscoverPlaylist(c echo.Context) error {
	client := spotifyR.GetClient(c)
	store := stmw.GetStore(c)
	queue := queuemw.GetUserQueue(c)
	discoverID, err := FindDiscoverPlaylist(client)
	log.Printf("%+v", discoverID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response, err := client.GetPlaylistTracksOpt("spotify", discoverID, nil, "items(track(album(images(url,height,width)),name,id,artists(name,id)))")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	responseJSON, err := json.Marshal(response.Tracks)
	err = json.Unmarshal(responseJSON, &SpotifyResponses)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	queueID := queue.ID
	log.Print(queueID)
	for _, track := range SpotifyResponses {
		SpotifyTracks = append(SpotifyTracks, trackR.Track{SpotifyTrack: track.Track})
	}
	updates := map[string]interface{}{}
	updates["TrackQueue"] = SpotifyTracks
	updates["trackqueue"] = SpotifyTracks
	err = store.UpdateByKey("queues", updates, "_id", queueID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, SpotifyTracks)
}

var SpotifyResponses []spotifyR.SpotifyResponse
var SpotifyTracks []trackR.Track
