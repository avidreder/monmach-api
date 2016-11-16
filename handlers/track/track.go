package track

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"github.com/labstack/echo"
	"gopkg.in/pg.v5"
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
	err = store.Create(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing track in the store
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
		if k == "ImageUrl" {
			payload["image_url"] = v[0]
		} else if k == "SpotifyID" {
			payload["spotify_id"] = v[0]
		} else if k == "Artists" {
			var array []string
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["artists"] = pg.Array(array)
		} else if k == "Genres" {
			var array []string
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["genres"] = pg.Array(array)
		} else if k == "Playlists" {
			var array []string
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["playlists"] = pg.Array(array)
		} else if k == "Features" {
			var array []float64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["features"] = pg.Array(array)
		} else if k == "TrackWhitelist" {
			var array []float64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["track_whitelist"] = pg.Array(array)
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

// Get retrieves an existing track in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	result := trackR.Track{ID: numID}
	store := stmw.GetStore(c)
	err = store.Get(&result)
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
	err := store.GetAll(&tracks, tableName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tracks)
}

// Delete deletes an existing track in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	track := trackR.Track{ID: numID}
	store := stmw.GetStore(c)
	err = store.Delete(&track)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Track %s deleted", id))
}
