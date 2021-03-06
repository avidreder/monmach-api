package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	"github.com/avidreder/monmach-api/resources/queue"

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
		err := store.GetByKey(user.ID, "queues", &userQueue, "userid", user.ID)
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
		client := spotifymw.GetClient(c)
		playlistOwner, err := spotifymw.FindPlaylistOwner(client, spotify.ID(playlistID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting playlist owner: %v", err))
		}
		tracks, err := spotifymw.TracksFromPlaylist(client, spotify.ID(playlistID), playlistOwner, user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting tracks: %v", err))
		}
		queueUpdates := structs.Map(queue.Queue{})
		queueID := bson.NewObjectId()
		queueUpdates["_id"] = queueID
		queueUpdates["userid"] = user.ID
		queueUpdates["ownerid"] = user.ID.Hex()
		queueUpdates["trackqueue"] = tracks
		queueUpdates["seedartists"] = make([]string, 0)
		queueUpdates["seedtracks"] = make([]string, 0)
		queueUpdates["listenedtracks"] = make([]string, 0)
		err = store.Create("queues", queueUpdates)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		userQueue := queue.Queue{}
		err = store.GetByKey(user.ID, "queues", &userQueue, "_id", queueID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve user queue")
		}
		log.Printf("user queue: %+v", userQueue)
		c.Set("queue", &userQueue)
		return h(c)
	}
}

func UpdateQueueTracks(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := usermw.GetUser(c)
		store := stmw.GetStore(c)
		currentQueue := GetUserQueue(c)
		newQueue := queue.Queue{}
		params, _ := c.FormParams()
		queueString := params["data"][0]
		log.Print(queueString)
		err := json.Unmarshal([]byte(queueString), &newQueue)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		queueUpdates := map[string]interface{}{}
		queueUpdates["trackqueue"] = newQueue.TrackQueue
		err = store.UpdateByKey(user.ID, "queues", queueUpdates, "_id", currentQueue.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return h(c)
	}
}
