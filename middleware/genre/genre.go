package genre

import (
	"encoding/json"
	"net/http"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	genreR "github.com/avidreder/monmach-api/resources/genre"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
)

// AddTrackToSeedTracks places a user into the contest
func AddTrackToSeedTracks(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		genre := genreR.Genre{}
		newTrack := trackR.Track{}
		params, _ := c.FormParams()
		trackString := params["data"][0]
		err := json.Unmarshal([]byte(trackString), &newTrack)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		newTrack.ID = bson.NewObjectId()
		err = store.GetByKey("genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, e := range genre.ListenedTracks {
			if e.SpotifyID == newTrack.SpotifyID {
				return echo.NewHTTPError(http.StatusInternalServerError, "Track already in playlist")
			}
		}
		newSeedTracks := append(genre.SeedTracks, newTrack)
		newListenedTracks := append(genre.ListenedTracks, newTrack)
		payload := map[string]interface{}{}
		payload["seedtracks"] = newSeedTracks
		payload["listenedtracks"] = newListenedTracks
		err = store.UpdateByKey("genres", payload, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// AddTrackToListened places a user into the contest
func AddTrackToListened(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		genre := genreR.Genre{}
		newTrack := trackR.Track{}
		params, _ := c.FormParams()
		trackString := params["data"][0]
		err := json.Unmarshal([]byte(trackString), &newTrack)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		newTrack.ID = bson.NewObjectId()
		err = store.GetByKey("genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, e := range genre.ListenedTracks {
			if e.SpotifyID == newTrack.SpotifyID {
				return echo.NewHTTPError(http.StatusInternalServerError, "Track already in playlist")
			}
		}
		newListenedTracks := append(genre.ListenedTracks, newTrack)
		payload := map[string]interface{}{}
		payload["listenedtracks"] = newListenedTracks
		err = store.UpdateByKey("genres", payload, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// GetUserGenres gets custom genres for a user
func GetUserGenres(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := usermw.GetUser(c)
		store := stmw.GetStore(c)
		genres := []genreR.Genre{}
		err := store.GetManyByKey("genres", &genres, "userid", user.ID.Hex())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		c.Set("result", genres)
		return h(c)
	}
}
