package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	authmw "github.com/avidreder/monmach-api/middleware/auth"
	stmw "github.com/avidreder/monmach-api/middleware/store"
	"github.com/avidreder/monmach-api/resources/store"
	userR "github.com/avidreder/monmach-api/resources/user"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"gopkg.in/pg.v5"
)

func init() {
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("monmach"))
}

// LogoutUser ends a user session
func LogoutUser(c echo.Context) error {
	sessionStore := authmw.GetStore(c)
	session, err := sessionStore.Get(c.Request(), "auth-session")
	log.Printf("authSession: %v", session)
	if err != nil {
		http.Redirect(c.Response().Writer(), c.Request(), "/", 302)
		return nil
	}
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response().Writer())
	session, err = gothic.Store.Get(c.Request(), "_gothic_session")
	log.Printf("gothicSession: %v", session)
	if err != nil {
		http.Redirect(c.Response().Writer(), c.Request(), "/", 302)
		return nil
	}
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response().Writer())
	http.Redirect(c.Response().Writer(), c.Request(), "/login", 302)
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
	provider := authmw.GetSpotifyProvider(c)
	q := c.Request().URL.Query()
	q.Add("provider", "spotify")
	c.Request().URL.RawQuery = q.Encode()
	goth.UseProviders(provider)
	gothic.BeginAuthHandler(c.Response().Writer(), c.Request())
	return nil
}

// FinishAuth finishes logging in the user
func FinishAuth(c echo.Context) error {
	store := stmw.GetStore(c)
	sessionStore := authmw.GetStore(c)
	q := c.Request().URL.Query()
	q.Add("provider", "spotify")
	c.Request().URL.RawQuery = q.Encode()
	response, err := gothic.CompleteUserAuth(c.Response().Writer(), c.Request())
	if err != nil {
		log.Printf("Could not log the user in: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not log the user in: %v", err))
	}
	user := userR.User{}
	string, _ := json.Marshal(response)
	err = json.Unmarshal(string, &user)
	if err != nil {
		log.Printf("Could not log the user in: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not log the user in: %v", err))
	}
	session, err := sessionStore.New(c.Request(), "auth-session")
	if err != nil {
		log.Printf("Could not log the user in: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not log the user in: %v", err))
	}
	session.Values["email"] = user.Email
	session.Save(c.Request(), c.Response().Writer())
	http.Redirect(c.Response().Writer(), c.Request(), "/", 302)
	go HandleUserLogin(user, store)
	return nil
}

func HandleUserLogin(user userR.User, store store.Store) {
	oldUser := userR.User{Email: user.Email}
	err := store.GetByEmail(&oldUser, user.Email)
	if err != nil {
		err = store.Create(&user)
		if err != nil {
			log.Printf("Error storing new user: %+v", user.Email)
			return
		}
		log.Printf("Stored new user: %+v", user.Email)
		return
	}
	values := map[string]interface{}{}
	values["name"] = user.Name
	values["spotify_token"] = user.SpotifyToken
	values["spotify_refresh_token"] = user.SpotifyRefreshToken
	values["spotify_id"] = user.SpotifyID
	values["track_blacklist"] = pg.Array(oldUser.TrackBlacklist)
	err = store.UpdateByEmail("users", user.Email, values)
	if err != nil {
		log.Printf("Error updating user: %+v", user.Email)
		return
	}
	log.Printf("Updated user: %+v", user.Email)
	return
}
