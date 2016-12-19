package crud

import (
	"errors"
	"net/http"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	genreR "github.com/avidreder/monmach-api/resources/genre"
	playlistR "github.com/avidreder/monmach-api/resources/playlist"
	queueR "github.com/avidreder/monmach-api/resources/queue"
	storeR "github.com/avidreder/monmach-api/resources/store"
	trackR "github.com/avidreder/monmach-api/resources/track"
	userR "github.com/avidreder/monmach-api/resources/user"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// Create inserts a new user into the store
func Create(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		table := c.Param("table")
		store := stmw.GetStore(c)
		payload := map[string]interface{}{}
		form, _ := c.FormParams()
		for k, v := range form {
			payload[k] = v[0]
		}
		switch table {
		case "genres":
			model := genreR.Genre{}
			payload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "playlists":
			model := playlistR.Playlist{}
			payload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "queues":
			model := queueR.Queue{}
			payload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "tracks":
			model := trackR.Track{}
			payload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "users":
			model := userR.User{}
			payload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
		}
	}
}

// Update inserts a new user into the store
func Update(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if !bson.IsObjectIdHex(id) {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.New("ID was not a valid ObjectID"))
		}
		table := c.Param("table")
		store := stmw.GetStore(c)
		payload := map[string]interface{}{}
		form, _ := c.FormParams()
		for k, v := range form {
			payload[k] = v[0]
		}
		switch table {
		case "genres":
			model := genreR.Genre{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "playlists":
			model := playlistR.Playlist{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "queues":
			model := queueR.Queue{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "tracks":
			model := trackR.Track{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "users":
			model := userR.User{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
		}
	}
}

// Get retrieves an existing genre in the store
func Get(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if !bson.IsObjectIdHex(id) {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.New("ID was not a valid ObjectID"))
		}
		table := c.Param("table")
		store := stmw.GetStore(c)
		switch table {
		case "genres":
			model := genreR.Genre{}
			err := store.GetByKey(table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "playlists":
			model := playlistR.Playlist{}
			err := store.GetByKey(table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "queues":
			model := queueR.Queue{}
			err := store.GetByKey(table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "tracks":
			model := trackR.Track{}
			err := store.GetByKey(table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "users":
			model := userR.User{}
			err := store.GetByKey(table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
		}
	}
}

// GetAll retrieves an existing genre in the store
func GetAll(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		table := c.Param("table")
		store := stmw.GetStore(c)
		switch table {
		case "genres":
			model := []genreR.Genre{}
			err := store.GetAll(table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "playlists":
			model := []playlistR.Playlist{}
			err := store.GetAll(table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "queues":
			model := []queueR.Queue{}
			err := store.GetAll(table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "tracks":
			model := []trackR.Track{}
			err := store.GetAll(table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "users":
			model := []userR.User{}
			err := store.GetAll(table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
		}
	}
}

// Get retrieves an existing genre in the store
func Delete(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if !bson.IsObjectIdHex(id) {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.New("ID was not a valid ObjectID"))
		}
		table := c.Param("table")
		store := stmw.GetStore(c)
		err := store.DeleteByKey(table, "_id", bson.ObjectIdHex(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}