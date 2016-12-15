package genre

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	genreR "github.com/avidreder/monmach-api/resources/genre"
	storeR "github.com/avidreder/monmach-api/resources/store"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

const tableName = "genres"

// Create inserts a new genre into the store
func Create(c echo.Context) error {
	var err error
	store := stmw.GetStore(c)
	payload := genreR.Genre{}
	userID := c.FormValue("UserID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "UserID is required")
	}
	queueID := c.FormValue("QueueID")
	if queueID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "QueueID is required")
	}
	name := c.FormValue("Name")
	if name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Name is required")
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
	if c.FormValue("SeedPlaylists") != "" {
		err = json.Unmarshal([]byte(c.FormValue("SeedPlaylists")), &payload.SeedPlaylists)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("TrackBlacklist") != "" {
		err = json.Unmarshal([]byte(c.FormValue("TrackBlacklist")), &payload.TrackBlacklist)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("TrackWhitelist") != "" {
		err = json.Unmarshal([]byte(c.FormValue("TrackWhitelist")), &payload.TrackWhitelist)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	payload.UserID = userID
	payload.QueueID = queueID
	payload.Description = c.FormValue("Description")
	payload.AvatarURL = c.FormValue("AvatarURL")
	payload.Name = name
	err = store.Create(tableName, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing genre in the store
func Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	store := stmw.GetStore(c)
	genre := genreR.Genre{}
	form, _ := c.FormParams()
	payload := map[string]interface{}{}
	for k, v := range form {
		payload[k] = v[0]
	}
	payload = storeR.ValidateInputs(genre, payload)
	err := store.UpdateByKey(tableName, payload, "_id", bson.ObjectIdHex(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, "Update was successful")
}

// Get retrieves an existing genre in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	result := genreR.Genre{}
	store := stmw.GetStore(c)
	err := store.GetByKey(tableName, &result, "_id", bson.ObjectIdHex(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Print(result)
	return c.JSON(http.StatusOK, result)
}

// GetAll retrieves all existing genres in the store
func GetAll(c echo.Context) error {
	var genres []genreR.Genre
	store := stmw.GetStore(c)
	err := store.GetAll(tableName, &genres)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, genres)
}

// Delete deletes an existing genre in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	store := stmw.GetStore(c)
	err := store.DeleteByKey(tableName, "_id", bson.ObjectIdHex(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Genre %s deleted", id))
}
