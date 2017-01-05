package queue

import (
	"fmt"
	"log"
	"net/http"

	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	"github.com/avidreder/monmach-api/resources/queue"
	spotifyR "github.com/avidreder/monmach-api/resources/spotify"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/zmb3/spotify"
	"gopkg.in/mgo.v2/bson"
)

// GetUserQueue retieves the user queue from the context
func GetUserQueue(c echo.Context) *queue.Queue {
	return c.Get("queue").(*queue.Queue)
}

// LoadUserQueue places a user into the contest
func LoadUserQueue(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := usermw.GetUser(c)
		userQueue := queue.Queue{}
		store := stmw.GetStore(c)
		err := store.GetByKey("queues", &userQueue, "userid", user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve user queue")
		}
		log.Printf("user queue: %+v", userQueue)
		c.Set("queue", &userQueue)
		return h(c)
	}
}

// QueueFromPlaylist places a user into the contest
func QueueFromPlaylist(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		playlistID := c.Param("playlist")
		if playlistID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid playlist ID")
		}
		user := usermw.GetUser(c)
		store := stmw.GetStore(c)
		client := spotifyR.GetClient(c)
		tracks, err := spotifymw.TracksFromPlaylist(client, spotify.ID(playlistID), user.SpotifyID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting tracks: %v", err))
		}
		queueUpdates := structs.Map(queue.Queue{})
		queueID := bson.NewObjectId()
		queueUpdates["_id"] = queueID
		queueUpdates["ID"] = queueID
		queueUpdates["userid"] = user.ID
		queueUpdates["UserID"] = user.ID
		queueUpdates["TrackQueue"] = tracks
		queueUpdates["trackqueue"] = tracks
		err = store.Create("queues", queueUpdates)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		userQueue := queue.Queue{}
		err = store.GetByKey("queues", &userQueue, "_id", queueID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve user queue")
		}
		log.Printf("user queue: %+v", userQueue)
		c.Set("queue", &userQueue)
		return h(c)
	}
}
