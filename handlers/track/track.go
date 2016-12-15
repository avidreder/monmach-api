package track

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	storeR "github.com/avidreder/monmach-api/resources/store"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"github.com/labstack/echo"
)

const tableName = "tracks"

// Create inserts a new track into the store
func Create(c echo.Context) error {
	store := stmw.GetStore(c)
	payload := trackR.Track{}
	rating := c.FormValue("Rating")
	numRating, err := strconv.ParseInt(rating, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "rating cannot be a string")
	}
	spotifyID := c.FormValue("SpotifyID")
	if spotifyID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "SpotifyID is required")
	}
	name := c.FormValue("Name")
	if name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Name is required")
	}
	artists := c.FormValue("Artists")
	if artists == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Artists is required")
	}
	err = json.Unmarshal([]byte(artists), &payload.Artists)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if c.FormValue("Features") != "" {
		err = json.Unmarshal([]byte(c.FormValue("Features")), &payload.Features)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("Genres") != "" {
		err = json.Unmarshal([]byte(c.FormValue("Genres")), &payload.Genres)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if c.FormValue("Playlists") != "" {
		err = json.Unmarshal([]byte(c.FormValue("Playlists")), &payload.Playlists)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	payload.Name = name
	payload.Rating = numRating
	payload.SpotifyID = spotifyID
	payload.ImageURL = c.FormValue("ImageURL")
	err = store.Create(tableName, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing track in the store
func Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	store := stmw.GetStore(c)
	track := trackR.Track{}
	form, _ := c.FormParams()
	payload := map[string]interface{}{}
	for k, v := range form {
		payload[k] = v[0]
	}
	payload = storeR.ValidateInputs(track, payload)
	err := store.UpdateByKey(tableName, payload, "_id", bson.ObjectIdHex(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, "Update was successful")
}

// Get retrieves an existing track in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
	}
	result := trackR.Track{}
	store := stmw.GetStore(c)
	err := store.GetByKey(tableName, &result, "_id", bson.ObjectIdHex(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Print(result)
	return c.JSON(http.StatusOK, result)
}

// GetAll retrieves all existing tracks in the store
func GetAll(c echo.Context) error {
	var tracks []trackR.Track
	store := stmw.GetStore(c)
	err := store.GetAll(tableName, &tracks)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tracks)
}

// Delete deletes an existing track in the store
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
	return c.JSON(http.StatusOK, fmt.Sprintf("Track %s deleted", id))
}
