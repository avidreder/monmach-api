package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	userR "github.com/avidreder/monmach-api/resources/user"

	"github.com/labstack/echo"
)

const tableName = "users"

// Create inserts a new user into the store
func Create(c echo.Context) error {
	store := stmw.GetStore(c)
	payload := userR.User{}
	name := c.FormValue("Name")
	if name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Name is required")
	}
	spotifyToken := c.FormValue("SpotifyToken")
	if spotifyToken == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Spotify Token is required")
	}
	spotifyRefreshToken := c.FormValue("SpotifyRefreshToken")
	if spotifyRefreshToken == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "SpotifyRefreshToken is required")
	}
	spotifyID := c.FormValue("SpotifyID")
	if spotifyID == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "SpotifyID is required")
	}
	email := c.FormValue("Email")
	if email == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "Email is required")
	}
	if c.FormValue("TrackBlacklist") != "" {
		err := json.Unmarshal([]byte(c.FormValue("TrackBlacklist")), &payload.TrackBlacklist)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	payload.Email = email
	payload.AvatarURL = c.FormValue("AvatarURL")
	payload.Name = name
	payload.SpotifyToken = spotifyToken
	payload.SpotifyRefreshToken = spotifyRefreshToken
	payload.SpotifyID = spotifyID
	err := store.Create(tableName, &payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payload)
}

// Update updates an existing user in the store
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

// Get retrieves an existing user in the store
func Get(c echo.Context) error {
	id := c.Param("id")
	numID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "id cannot be a string")
	}
	result := userR.User{ID: numID}
	store := stmw.GetStore(c)
	err = store.GetByKey(tableName, &result, "_id", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	log.Print(result)
	return c.JSON(http.StatusOK, result)
}

// GetAll retrieves all existing users in the store
func GetAll(c echo.Context) error {
	var users []userR.User
	store := stmw.GetStore(c)
	err := store.GetAll(tableName, &users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// Delete deletes an existing user in the store
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
	return c.JSON(http.StatusOK, fmt.Sprintf("User %s deleted", id))
}
