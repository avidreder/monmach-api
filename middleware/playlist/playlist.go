package queue

import (
	"fmt"
	"net/http"

	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"github.com/labstack/echo"
	"github.com/zmb3/spotify"
)

// GetTracks retieves the user queue from the context
func GetTracks(c echo.Context) *[]trackR.Track {
	return c.Get("tracks").(*[]trackR.Track)
}

// TracksFromPlaylist places a user into the contest
func TracksFromPlaylist(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		playlistID := c.Param("playlist")
		if playlistID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid playlist ID")
		}
		client := spotifymw.GetClient(c)
		playlistOwner, err := spotifymw.FindPlaylistOwner(client, spotify.ID(playlistID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting playlist owner: %v", err))
		}
		tracks, err := spotifymw.TracksFromPlaylist(client, spotify.ID(playlistID), playlistOwner)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting tracks: %v", err))
		}
		c.Set("tracks", &tracks)
		return h(c)
	}
}
