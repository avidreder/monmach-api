package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	authmw "github.com/avidreder/monmach-api/middleware/auth"
	spotifymw "github.com/avidreder/monmach-api/middleware/spotify"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	configR "github.com/avidreder/monmach-api/resources/config"
	"github.com/avidreder/monmach-api/resources/queue"
	"github.com/avidreder/monmach-api/resources/store"
	trackR "github.com/avidreder/monmach-api/resources/track"
	userR "github.com/avidreder/monmach-api/resources/user"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// LogoutUser ends a user session
func LogoutUser(c echo.Context) error {
	sessionStore := authmw.GetStore(c)
	session, err := sessionStore.Get(c.Request(), "auth-session")
	if err != nil {
		c.Redirect(302, configR.CurrentConfig.ClientAddress)
		return nil
	}
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response().Writer())
	c.Redirect(302, configR.CurrentConfig.ClientAddress)
	return nil
}

// GetUser ends a user session
func GetUser(c echo.Context) error {
	sessionStore := authmw.GetStore(c)
	session, err := sessionStore.Get(c.Request(), "auth-session")
	payload := struct {
		LoggedIn bool
		Email    string
	}{}
	if session.IsNew || err != nil {
		return c.JSON(404, payload)
	}
	payload.Email = session.Values["email"].(string)
	payload.LoggedIn = true
	return c.JSON(200, payload)
}

// StartAuth begins authorization
func StartAuth(c echo.Context) error {
	auth := spotifymw.GetAuthenticator(c)
	c.Redirect(301, auth.AuthURL("state"))
	return nil
}

// FinishAuth finishes logging in the user
func FinishAuth(c echo.Context) error {
	store := stmw.GetStore(c)
	sessionStore := authmw.GetStore(c)
	auth := spotifymw.GetAuthenticator(c)
	token, err := auth.Token("state", c.Request())
	if err != nil {
		log.Printf("Could not get token from Spotify: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not get token from Spotify"))
	}
	client := auth.NewClient(token)
	spotifyUser, err := client.CurrentUser()
	if err != nil {
		log.Printf("Could not get user from Spotify: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not get user from Spotify"))
	}
	user := userR.User{}
	user.Token = *token
	user.Email = spotifyUser.Email
	user.Name = spotifyUser.DisplayName
	user.SpotifyID = spotifyUser.ID
	session, err := sessionStore.New(c.Request(), "auth-session")
	if err != nil {
		session.Values["email"] = user.Email
		session.Save(c.Request(), c.Response().Writer())
		log.Printf("Creating new session: %v", err)
	}
	if session.IsNew {
		session.Values["email"] = user.Email
		session.Save(c.Request(), c.Response().Writer())
	}
	c.Redirect(302, configR.CurrentConfig.ClientAddress)
	go HandleUserLogin(user, store)
	return nil
}

// HandleUserLogin creates or updates a user record, and it's associated queue
func HandleUserLogin(user userR.User, store store.Store) {
	oldUser := userR.User{}
	err := store.GetByKey("users", &oldUser, "Email", user.Email)
	if err != nil && oldUser.Email != user.Email {
		updates := structs.Map(user)
		updates["Created"] = time.Now()
		updates["Updated"] = time.Now()
		id := bson.NewObjectId()
		updates["ID"] = id
		updates["_id"] = id
		err = store.Create("users", updates)
		if err != nil {
			log.Printf("Error storing new user: %+v, %+v", user.Email, err)
			return
		}
		log.Printf("Stored new user: %+v", user.Email)
		queueUpdates := structs.Map(queue.Queue{})
		queueID := bson.NewObjectId()
		queueUpdates["_id"] = queueID
		queueUpdates["ID"] = queueID
		queueUpdates["userid"] = id
		queueUpdates["UserID"] = id
		queueUpdates["TrackQueue"] = make([]trackR.Track, 0)
		queueUpdates["SeedArtists"] = make([]string, 0)
		queueUpdates["SeedTracks"] = make([]string, 0)
		queueUpdates["ListenedTracks"] = make([]string, 0)
		queueUpdates["trackqueue"] = make([]trackR.Track, 0)
		queueUpdates["seedartists"] = make([]string, 0)
		queueUpdates["seedtracks"] = make([]string, 0)
		queueUpdates["listenedtracks"] = make([]string, 0)
		err = store.Create("queues", queueUpdates)
		if err != nil {
			log.Printf("Error creating user queue: %+v, %+v", id, err)
			return
		}
		log.Printf("Created user queue: %+v", id)
		return
	}
	updates := map[string]interface{}{}
	updates["Token"] = user.Token
	updates["Updated"] = time.Now()
	err = store.UpdateByKey("users", updates, "Email", user.Email)
	if err != nil {
		log.Printf("Error updating user: %+v, %+v", user.Email, err)
		return
	}
	log.Printf("Updated user: %+v", user.Email)
	return
}
