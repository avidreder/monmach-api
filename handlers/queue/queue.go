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
	payload.UserID = userID
	payload.MaxSize = numMaxSize
	payload.Name = name
	err = store.Create(tableName, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing queue in the store
func Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	store := stmw.GetStore(c)
	form, _ := c.FormParams()
	payload := map[string]interface{}{}
	for k, v := range form {
		payload[k] = v
	}
	err := store.UpdateByKey(tableName, payload, "_id", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, "Update was successful")
}

// Get retrieves an existing queue in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	result := queueR.Queue{}
	store := stmw.GetStore(c)
	err := store.GetByKey(tableName, &result, "_id", id)
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
	err := store.GetAll(tableName, &queues)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, queues)
}

// Delete deletes an existing queue in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	store := stmw.GetStore(c)
	err := store.DeleteByKey(tableName, "_id", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Queue %s deleted", id))
}
