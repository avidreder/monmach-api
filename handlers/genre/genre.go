package genre

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	genreR "github.com/avidreder/monmach-api/resources/genre"

	"github.com/labstack/echo"
	"gopkg.in/pg.v5"
)

const tableName = "genres"

// Create inserts a new genre into the store
func Create(c echo.Context) error {
	store := stmw.GetStore(c)
	payload := genreR.Genre{}
	userID := c.FormValue("UserID")
	if userID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "UserID is required")
	}
	numUserID, err := strconv.ParseInt(userID, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	queueID := c.FormValue("QueueID")
	if queueID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "QueueID is required")
	}
	numQueueID, err := strconv.ParseInt(queueID, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
	payload.UserID = numUserID
	payload.QueueID = numQueueID
	payload.Description = c.FormValue("Description")
	payload.AvatarURL = c.FormValue("AvatarURL")
	payload.Name = name
	err = store.Create(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing genre in the store
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
		} else if k == "QueueID" {
			payload["queue_id"] = v[0]
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
		} else if k == "TrackWhitelist" {
			var array []int64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["track_whitelist"] = pg.Array(array)
		} else if k == "TrackBlacklist" {
			var array []int64
			err = json.Unmarshal([]byte(v[0]), &array)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			payload["track_blacklist"] = pg.Array(array)
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

// Get retrieves an existing genre in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	result := genreR.Genre{ID: numID}
	store := stmw.GetStore(c)
	err = store.Get(&result)
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
	err := store.GetAll(&genres, tableName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, genres)
}

// Delete deletes an existing genre in the store
func Delete(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	genre := genreR.Genre{ID: numID}
	store := stmw.GetStore(c)
	err = store.Delete(&genre)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("Genre %s deleted", id))
}
