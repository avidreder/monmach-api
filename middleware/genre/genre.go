package genre

import (
	"encoding/json"
	"log"
	"net/http"

	stmw "github.com/avidreder/monmach-api/middleware/store"
	usermw "github.com/avidreder/monmach-api/middleware/user"
	genreR "github.com/avidreder/monmach-api/resources/genre"
	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
)

// AddTrack places a user into the contest
func AddTrack(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		user := usermw.GetUser(c)
		genre := genreR.Genre{}
		newTrack := trackR.Track{}
		params, _ := c.FormParams()
		trackString := params["data"][0]
		err := json.Unmarshal([]byte(trackString), &newTrack)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		newTrack.ID = bson.NewObjectId()
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, e := range genre.TrackList {
			if e.SpotifyID == newTrack.SpotifyID {
				return echo.NewHTTPError(http.StatusInternalServerError, "Track already in track list")
			}
		}
		addToListened := true
		for _, e := range genre.ListenedTracks {
			if e.SpotifyID == newTrack.SpotifyID {
				addToListened = false
			}
		}
		newTrack.OwnerID = user.ID
		newTrackList := append(genre.TrackList, newTrack)
		payload := map[string]interface{}{}
		payload["tracklist"] = newTrackList
		if addToListened {
			newListenedTracks := append(genre.ListenedTracks, newTrack)
			payload["listenedtracks"] = newListenedTracks
		}
		err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// RemoveTrack places a user into the contest
func RemoveTrack(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		user := usermw.GetUser(c)
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
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for k, v := range genre.TrackList {
			if v.SpotifyID == newTrack.SpotifyID {
				newTrackList := append(genre.TrackList[:k], genre.TrackList[k+1:]...)
				payload := map[string]interface{}{}
				payload["tracklist"] = newTrackList
				err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				return h(c)
			}
		}
		return h(c)
	}
}

// AddTrackToSeedTracks places a user into the contest
func AddTrackToSeedTracks(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		user := usermw.GetUser(c)
		genre := genreR.Genre{}
		newTrack := trackR.Track{}
		params, _ := c.FormParams()
		trackString := params["data"][0]
		err := json.Unmarshal([]byte(trackString), &newTrack)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		newTrack.ID = bson.NewObjectId()
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, e := range genre.SeedTracks {
			if e.SpotifyID == newTrack.SpotifyID {
				return echo.NewHTTPError(http.StatusInternalServerError, "Track already in seeds")
			}
		}
		addToListened := true
		for _, e := range genre.ListenedTracks {
			if e.SpotifyID == newTrack.SpotifyID {
				addToListened = false
			}
		}
		newTrack.OwnerID = user.ID
		newSeedTracks := append(genre.SeedTracks, newTrack)
		payload := map[string]interface{}{}
		payload["seedtracks"] = newSeedTracks
		if addToListened {
			newListenedTracks := append(genre.ListenedTracks, newTrack)
			payload["listenedtracks"] = newListenedTracks
		}
		err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// AddGenreToSeedGenres places a user into the contest
func AddGenreToSeedGenres(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		user := usermw.GetUser(c)
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		genre := genreR.Genre{}
		params, _ := c.FormParams()
		genreString := params["data"][0]
		log.Printf("genre is: %v", genreString)
		err := store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, e := range genre.SeedGenres {
			if e == genreString {
				return echo.NewHTTPError(http.StatusInternalServerError, "Genre already in seeds")
			}
		}
		newSeedGenres := append(genre.SeedGenres, genreString)
		payload := map[string]interface{}{}
		payload["seedgenres"] = newSeedGenres
		err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// AddArtistToSeedArtists places a user into the contest
func AddArtistToSeedArtists(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		user := usermw.GetUser(c)
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		genre := genreR.Genre{}
		newArtist := spotifyR.SpotifyArtist{}
		params, _ := c.FormParams()
		artistString := params["data"][0]
		err := json.Unmarshal([]byte(artistString), &newArtist)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, e := range genre.SeedArtists {
			if e.SpotifyID == newArtist.SpotifyID {
				return echo.NewHTTPError(http.StatusInternalServerError, "Artist already in genre")
			}
		}
		newSeedArtists := append(genre.SeedArtists, newArtist)
		payload := map[string]interface{}{}
		payload["seedartists"] = newSeedArtists
		err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}

// RemoveArtistFromSeedArtists places a user into the contest
func RemoveArtistFromSeedArtists(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		user := usermw.GetUser(c)
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		genre := genreR.Genre{}
		newArtist := spotifyR.SpotifyArtist{}
		params, _ := c.FormParams()
		artistString := params["data"][0]
		err := json.Unmarshal([]byte(artistString), &newArtist)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for k, v := range genre.SeedArtists {
			if v.SpotifyID == newArtist.SpotifyID {
				newSeedArtists := append(genre.SeedArtists[:k], genre.SeedArtists[k+1:]...)
				payload := map[string]interface{}{}
				payload["seedartists"] = newSeedArtists
				err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				return h(c)
			}
		}
		return h(c)
	}
}

// RemoveTrackFromSeedTracks places a user into the contest
func RemoveTrackFromSeedTracks(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		user := usermw.GetUser(c)
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
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for k, v := range genre.SeedTracks {
			if v.SpotifyID == newTrack.SpotifyID {
				newSeedTracks := append(genre.SeedTracks[:k], genre.SeedTracks[k+1:]...)
				payload := map[string]interface{}{}
				payload["seedtracks"] = newSeedTracks
				err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				return h(c)
			}
		}
		return h(c)
	}
}

// RemoveGenreFromSeedGenres places a user into the contest
func RemoveGenreFromSeedGenres(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		genreID := c.Param("id")
		if genreID == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, "Not a valid genre ID")
		}
		user := usermw.GetUser(c)
		bsonID := bson.ObjectIdHex(genreID)
		store := stmw.GetStore(c)
		genre := genreR.Genre{}
		params, _ := c.FormParams()
		genreString := params["data"][0]
		err := store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for k, v := range genre.SeedGenres {
			if genreString == v {
				newSeedGenres := append(genre.SeedGenres[:k], genre.SeedGenres[k+1:]...)
				payload := map[string]interface{}{}
				payload["seedgenres"] = newSeedGenres
				err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				return h(c)
			}
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
		user := usermw.GetUser(c)
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
		err = store.GetByKey(user.ID, "genres", &genre, "_id", bsonID)
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
		err = store.UpdateByKey(user.ID, "genres", payload, "_id", bsonID)
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
		err := store.GetManyByKey(user.ID, "genres", &genres, "userid", user.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		c.Set("result", genres)
		return h(c)
	}
}

// CreateNewGenre gets custom genres for a user
func CreateNewGenre(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		trackParams := struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}{}
		user := usermw.GetUser(c)
		store := stmw.GetStore(c)
		params, _ := c.FormParams()
		paramString := params["data"][0]
		err := json.Unmarshal([]byte(paramString), &trackParams)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		count, err := store.CountByQuery(user.ID, "genres", "name", trackParams.Name)
		log.Print("name")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if count > 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Genre Name already taken")
		}
		fields := map[string]interface{}{}
		fields["name"] = trackParams.Name
		fields["description"] = trackParams.Description
		fields["userid"] = user.ID
		fields["ownerid"] = user.ID
		fields["trackqueue"] = make([]trackR.Track, 0)
		fields["seedartists"] = make([]spotifyR.SpotifyArtist, 0)
		fields["seedtracks"] = make([]trackR.Track, 0)
		fields["tracklist"] = make([]trackR.Track, 0)
		fields["seedgenres"] = make([]string, 0)
		fields["listenedtracks"] = make([]string, 0)
		err = store.Create("genres", fields)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return h(c)
	}
}
