package spotify

import (
	"log"
	"net/http"

	"github.com/avidreder/monmach-api/resources/spotify"

	"github.com/labstack/echo"
)

// UserPlaylists gets a user's spotify playlists
func UserPlaylists(c echo.Context) error {
	client := spotify.GetClient(c)
	log.Printf("The client: %+v", client)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, playlists)
}

// DiscoverPlaylist gets a user's spotify discover playlist
func DiscoverPlaylist(c echo.Context) error {
	client := spotify.GetClient(c)
	log.Printf("The client: %+v", client)
	tracks, err := client.GetPlaylistTracks("spotifydiscover", "5yjxqPpY8Ch9knz2rGW0CH")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tracks)
}
