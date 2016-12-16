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

// var singleTypeMap = map[string]interface{}{}
// var pluralTypeMap = map[string]interface{}{}
//
// func init() {
// 	singleTypeMap["genres"] = genreR.Genre{}
// 	singleTypeMap["playlists"] = playlistR.Playlist{}
// 	singleTypeMap["queue"] = queueR.Queue{}
// 	singleTypeMap["user"] = userR.User{}
// 	singleTypeMap["track"] = trackR.Track{}
//
// 	pluralTypeMap["genres"] = []genreR.Genre{}
// 	pluralTypeMap["playlists"] = []playlistR.Playlist{}
// 	pluralTypeMap["queue"] = []queueR.Queue{}
// 	pluralTypeMap["user"] = []userR.User{}
// 	pluralTypeMap["track"] = []trackR.Track{}
// }

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

// // ParseParams inserts a new user into the store
// func ParseParams(h echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		table := c.Param("table")
// 		c.Set("table", table)
// 		switch table {
// 		case "genres":
// 			c.Set("model", genreR.Genre{})
// 		case "playlists":
// 			c.Set("model", playlistR.Playlist{})
// 		case "queues":
// 			c.Set("model", queueR.Queue{})
// 		case "tracks":
// 			c.Set("model", trackR.Track{})
// 		case "users":
// 			c.Set("model", userR.User{})
// 		default:
// 			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
// 		}
// 		return h(c)
// 	}
// }
//
// // ParseParams inserts a new user into the store
// func ParseParamsPlural(h echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		table := c.Param("table")
// 		c.Set("table", table)
// 		switch table {
// 		case "genres":
// 			c.Set("model", []genreR.Genre{})
// 		case "playlists":
// 			c.Set("model", []playlistR.Playlist{})
// 		case "queues":
// 			c.Set("model", []queueR.Queue{})
// 		case "tracks":
// 			c.Set("model", []trackR.Track{})
// 		case "users":
// 			c.Set("model", []userR.User{})
// 		default:
// 			return echo.NewHTTPError(http.StatusInternalServerError, "Not found")
// 		}
// 		return h(c)
// 	}
// }

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
		if id == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
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
		if id == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "id is required")
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
