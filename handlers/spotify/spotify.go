package spotify

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	spotifyR "github.com/avidreder/monmach-api/resources/spotify"

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

// DiscoverPlaylist gets a user's spotify discover playlist
func DiscoverPlaylist(c echo.Context) error {
	client := spotifyR.GetClient(c)
	discoverID, err := FindDiscoverPlaylist(client)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response, err := client.GetPlaylistTracksOpt("spotifydiscover", discoverID, nil, "items(track(album(images(url,height,width)),name,id,artists(name,id)))")
	log.Printf("%+v", response)
	responseJSON, err := json.Marshal(response.Tracks)
	err = json.Unmarshal(responseJSON, &SpotifyTracks)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, SpotifyTracks)
}

var SpotifyTracks []spotifyR.SpotifyTrack
