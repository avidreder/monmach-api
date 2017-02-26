package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	usermw "github.com/avidreder/monmach-api/middleware/user"
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
		user := usermw.GetUser(c)
		playlistID := c.Param("playlist")
		if playlistID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid playlist ID")
		}
		client := spotifymw.GetClient(c)
		playlistOwner, err := spotifymw.FindPlaylistOwner(client, spotify.ID(playlistID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting playlist owner: %v", err))
		}
		tracks, err := spotifymw.TracksFromPlaylist(client, spotify.ID(playlistID), playlistOwner, user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting tracks: %v", err))
		}
		c.Set("tracks", &tracks)
		return h(c)
	}
}

// RecommendedTracks places a user into the contest
func RecommendedTracks(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := spotifymw.GetClient(c)
		user := usermw.GetUser(c)
		params, _ := c.FormParams()
		postParams := spotifymw.RecommendedTrackParams{}
		dataString := params["data"][0]
		log.Printf("string: %+v", dataString)
		err := json.Unmarshal([]byte(dataString), &postParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		log.Printf("rec params: %+v", postParams)
		tracks, err := spotifymw.RecommendedTracks(client, postParams, user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting tracks: %v", err))
		}
		c.Set("tracks", &tracks)
		return h(c)
	}
}
