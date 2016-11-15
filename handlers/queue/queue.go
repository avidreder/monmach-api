package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	queueR "github.com/avidreder/monmach-api/resources/queue"

	"github.com/labstack/echo"
	"gopkg.in/pg.v5"
)

const tableName = "queues"

// Create inserts a new queue into the store
func Create(c echo.Context) error {
	store := stmw.GetStore(c)
	payload := queueR.Queue{}
	userID := c.FormValue("UserID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "UserID is required")
	}
	numUserID, err := strconv.ParseInt(userID, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	maxSize := c.FormValue("MaxSize")
	if maxSize == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "MaxSize is required")
	}
	numMaxSize, err := strconv.ParseInt(maxSize, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	name := c.FormValue("Name")
	if name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Name is required")
	}
	if c.FormValue("TrackQueue") != "" {
		err = json.Unmarshal([]byte(c.FormValue("TrackQueue")), &payload.TrackQueue)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("SeedArtists") != "" {
		err = json.Unmarshal([]byte(c.FormValue("SeedArtists")), &payload.SeedArtists)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("SeedTracks") != "" {
		err = json.Unmarshal([]byte(c.FormValue("SeedTracks")), &payload.SeedTracks)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("ListenedTracks") != "" {
		err = json.Unmarshal([]byte(c.FormValue("ListenedTracks")), &payload.ListenedTracks)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	payload.UserID = numUserID
	payload.MaxSize = numMaxSize
	payload.Name = name
	err = store.Create(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing queue in the store
func Update(c echo.Context) error {
	id := c.Param("id")
	numId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	store := stmw.GetStore(c)
	form, _ := c.FormParams()
	payload := map[string]interface{}{}
	for k, v := range form {
		if k == "UserID" {
			payload["user_id"] = v[0]
		} else if k == "MaxSize" {
			payload["max_size"] = v[0]
		} else if k == "TrackQueue" {
			var array []int64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["track_queue"] = pg.Array(array)
		} else if k == "SeedTracks" {
			var array []int64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["seed_tracks"] = pg.Array(array)
		} else if k == "SeedArtists" {
			var array []int64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["seed_artists"] = pg.Array(array)
		} else if k == "ListenedTracks" {
			var array []int64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["listened_tracks"] = pg.Array(array)
		} else {
			payload[k] = v[0]
		}
	}
	err = store.Update(tableName, numId, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, "Update was successful")
}

// Get retrieves an existing queue in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	result := queueR.Queue{ID: numID}
	store := stmw.GetStore(c)
	err = store.Get(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Print(result)
	return c.JSON(http.StatusOK, result)
}

// GetAll retrieves all existing queues in the store
func GetAll(c echo.Context) error {
	var queues []queueR.Queue
	store := stmw.GetStore(c)
	err := store.GetAll(&queues, tableName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, queues)
}

// Delete deletes an existing queue in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	queue := queueR.Queue{ID: numID}
	store := stmw.GetStore(c)
	err = store.Delete(&queue)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Queue %s deleted", id))
}
