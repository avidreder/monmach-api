package playlist

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	playlistR "github.com/avidreder/monmach-api/resources/playlist"

	"github.com/labstack/echo"
)

const tableName = "playlists"

// Create inserts a new playlist into the store
func Create(c echo.Context) error {
	store := stmw.GetStore(c)
	payload := playlistR.Playlist{}
	userID := c.FormValue("UserID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "UserID is required")
	}
	tracks := c.FormValue("Tracks")
	if tracks == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Tracks is required")
	}
	name := c.FormValue("Name")
	if userID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Name is required")
	}
	numID, err := strconv.ParseInt(userID, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	var array []int64
	err = json.Unmarshal([]byte(tracks), &array)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	payload.UserID = numID
	payload.Tracks = array
	payload.Name = name
	payload.SpotifyID = c.FormValue("SpotifyID")
	err = store.Create(tableName, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing playlist in the store
func Update(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	store := stmw.GetStore(c)
	form, _ := c.FormParams()
	payload := map[string]interface{}{}
	for k, v := range form {
		payload[k] = v
	}
	err = store.UpdateByKey(tableName, payload, "_id", numID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, "Update was successful")
}

// Get retrieves an existing playlist in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	result := playlistR.Playlist{ID: numID}
	store := stmw.GetStore(c)
	err = store.GetByKey(tableName, &result, "_id", numID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Print(result)
	return c.JSON(http.StatusOK, result)
}

// GetAll retrieves all existing playlists in the store
func GetAll(c echo.Context) error {
	var playlists []playlistR.Playlist
	store := stmw.GetStore(c)
	err := store.GetAll(tableName, &playlists)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, playlists)
}

// Delete deletes an existing playlist in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	store := stmw.GetStore(c)
	err = store.DeleteByKey(tableName, "_id", numID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Playlist %s deleted", id))
}
