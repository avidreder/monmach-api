package crud

import (
	"errors"
	"log"
	"net/http"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
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
		user := usermw.GetUser(c)
		payload := map[string]interface{}{}
		form, _ := c.FormParams()
		for k, v := range form {
			payload[k] = v[0]
		}
		switch table {
		case "genres":
			model := genreR.Genre{}
			createPayload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, createPayload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "playlists":
			model := playlistR.Playlist{}
			createPayload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, createPayload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "queues":
			model := queueR.Queue{}
			createPayload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, createPayload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "tracks":
			model := trackR.Track{}
			createPayload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			processed, err := trackR.AlreadyProcessed(user.ID, createPayload["spotifyid"].(string))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			log.Printf("Processed: %+v", processed)
			err = store.Create(table, createPayload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "users":
			model := userR.User{}
			createPayload, err := storeR.ValidateRequired(model, payload)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			err = store.Create(table, createPayload)
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
		user := usermw.GetUser(c)
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
			err := store.UpdateByKey(user.ID, table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "playlists":
			model := playlistR.Playlist{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(user.ID, table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "queues":
			model := queueR.Queue{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(user.ID, table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "tracks":
			model := trackR.Track{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(user.ID, table, payload, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return h(c)
		case "users":
			model := userR.User{}
			payload = storeR.ValidateInputs(model, payload)
			err := store.UpdateByKey(user.ID, table, payload, "_id", bson.ObjectIdHex(id))
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
		user := usermw.GetUser(c)
		id := c.Param("id")
		if !bson.IsObjectIdHex(id) {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.New("ID was not a valid ObjectID"))
		}
		table := c.Param("table")
		store := stmw.GetStore(c)
		switch table {
		case "genres":
			model := genreR.Genre{}
			err := store.GetByKey(user.ID, table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "playlists":
			model := playlistR.Playlist{}
			err := store.GetByKey(user.ID, table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "queues":
			model := queueR.Queue{}
			err := store.GetByKey(user.ID, table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "tracks":
			model := trackR.Track{}
			err := store.GetByKey(user.ID, table, &model, "_id", bson.ObjectIdHex(id))
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "users":
			model := userR.User{}
			err := store.GetByKey(user.ID, table, &model, "_id", bson.ObjectIdHex(id))
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
		user := usermw.GetUser(c)
		table := c.Param("table")
		store := stmw.GetStore(c)
		switch table {
		case "genres":
			model := []genreR.Genre{}
			err := store.GetAll(user.ID, table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "playlists":
			model := []playlistR.Playlist{}
			err := store.GetAll(user.ID, table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "queues":
			model := []queueR.Queue{}
			err := store.GetAll(user.ID, table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "tracks":
			model := []trackR.Track{}
			err := store.GetAll(user.ID, table, &model)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Set("result", model)
			return h(c)
		case "users":
			model := []userR.User{}
			err := store.GetAll(user.ID, table, &model)
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
		user := usermw.GetUser(c)
		table := c.Param("table")
		store := stmw.GetStore(c)
		err := store.DeleteByKey(user.ID, table, "_id", bson.ObjectIdHex(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}
