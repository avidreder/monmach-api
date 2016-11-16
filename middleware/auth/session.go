package auth

import (
	"fmt"
	"net/http"

	"github.com/avidreder/monmach-api/resources/auth"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
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

// GetStore retieves a data store from the context
func GetStore(c echo.Context) *sessions.FilesystemStore {
	return c.Get("sessionStore").(*sessions.FilesystemStore)
}

// CheckLogin checks for a valid session
func CheckLogin(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionStore := GetStore(c)
		session, err := sessionStore.Get(c.Request(), "auth-session")
		if session.IsNew || err != nil {
			return c.Redirect(302, "/login")
		}
		return h(c)
	}
}
