package spotify

import (
	"log"
	"net/http"

	queuemw "github.com/avidreder/monmach-api/middleware/queue"
	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	stmw "github.com/avidreder/monmach-api/middleware/store"

	"github.com/labstack/echo"
)

// UserPlaylists gets a user's spotify playlists
func UserPlaylists(c echo.Context) error {
	client := spotifymw.GetClient(c)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, playlists)
}

// DiscoverPlaylist gets and stores a user's spotify discover playlist
func DiscoverPlaylist(c echo.Context) error {
	client := spotifymw.GetClient(c)
	store := stmw.GetStore(c)
	queue := queuemw.GetUserQueue(c)
	discoverID, err := spotifymw.FindDiscoverPlaylist(client)
	log.Printf("%+v", discoverID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	tracks, err := spotifymw.TracksFromPlaylist(client, discoverID, "spotify")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	queueID := queue.ID
	log.Print(queueID)
	updates := map[string]interface{}{}
	updates["TrackQueue"] = tracks
	updates["trackqueue"] = tracks
	err = store.UpdateByKey("queues", updates, "_id", queueID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tracks)
}
