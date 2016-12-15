package crud

import (
	"fmt"
	"net/http"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	"github.com/avidreder/monmach-api/resources/auth"
	genreR "github.com/avidreder/monmach-api/resources/genre"
	playlistR "github.com/avidreder/monmach-api/resources/playlist"
	queueR "github.com/avidreder/monmach-api/resources/queue"
	storeR "github.com/avidreder/monmach-api/resources/store"
	trackR "github.com/avidreder/monmach-api/resources/track"
	userR "github.com/avidreder/monmach-api/resources/user"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// LoadStore places a data store in the context for later use
func LoadStore(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionStore, err := auth.Get()
		if err != nil {
			errorMessage := fmt.Sprintf("Could not load session store into context: %s", err)
			return echo.NewHTTPError(http.StatusUnauthorized, errorMessage)
		}
		c.Set("sessionStore", sessionStore)
		return h(c)
	}
}

// ParseParams inserts a new user into the store
func ParseParams(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		table := c.Param("table")
		c.Set("table", table)
		switch table {
		case "genre":
			c.Set("model", genreR.Genre{})
		case "playlist":
			c.Set("model", playlistR.Playlist{})
		case "queue":
			c.Set("model", queueR.Queue{})
		case "track":
			c.Set("model", trackR.Track{})
		case "user":
			c.Set("model", userR.User{})
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
		}
		return h(c)
	}
}

// Create inserts a new user into the store
func Create(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		table := c.Get("table").(string)
		model := c.Get("model")
		store := stmw.GetStore(c)
		payload := map[string]interface{}{}
		form, _ := c.FormParams()
		for k, v := range form {
			payload[k] = v[0]
		}
		payload, err := storeR.ValidateRequired(model, payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		err = store.Create(table, &payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// Update inserts a new user into the store
func Update(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
		}
		table := c.Get("table").(string)
		model := c.Get("model")
		store := stmw.GetStore(c)
		payload := map[string]interface{}{}
		form, _ := c.FormParams()
		for k, v := range form {
			payload[k] = v[0]
		}
		payload = storeR.ValidateInputs(model, payload)
		err := store.UpdateByKey(table, payload, "_id", bson.ObjectIdHex(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// Get retrieves an existing genre in the store
func Get(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
		}
		table := c.Get("table").(string)
		model := c.Get("model")
		store := stmw.GetStore(c)
		err := store.GetByKey(table, &model, "_id", bson.ObjectIdHex(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		c.Set("result", model)
		return h(c)
	}
}

// GetAll retrieves an existing genre in the store
func GetAll(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		table := c.Get("table").(string)
		model := c.Get("model")
		store := stmw.GetStore(c)
		err := store.GetAll(table, &model)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		c.Set("result", model)
		return h(c)
	}
}

// Get retrieves an existing genre in the store
func Delete(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
		}
		table := c.Get("table").(string)
		store := stmw.GetStore(c)
		err := store.DeleteByKey(table, "_id", bson.ObjectIdHex(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}
