package spotify

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/avidreder/monmach-api/resources/spotify"

	"github.com/labstack/echo"
)

// UserPlaylists gets a user's spotify playlists
func UserPlaylists(c echo.Context) error {
	client := spotify.GetClient(c)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, playlists)
}

// DiscoverPlaylist gets a user's spotify discover playlist
func DiscoverPlaylist(c echo.Context) error {
	client := spotify.GetClient(c)
	response, err := client.GetPlaylistTracksOpt("spotifydiscover", "5yjxqPpY8Ch9knz2rGW0CH", nil, "items(track(images(url),name,id,artists(name,id)))")
	responseJSON, err := json.Marshal(response.Tracks)
	log.Printf("tracks: %+v", string(responseJSON))
	err = json.Unmarshal(responseJSON, &SpotifyTracks)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, SpotifyTracks)
}

var SpotifyTracks []spotify.SpotifyTrack
