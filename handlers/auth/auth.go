package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"

	authmw "github.com/avidreder/monmach-api/middleware/auth"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func init() {
	gothic.Store = sessions.NewFilesystemStore(os.TempDir(), []byte("monmach"))
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
	q := c.Request().URL.Query()
	q.Add("provider", "spotify")
	c.Request().URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Response().Writer(), c.Request())
	if err != nil {
		log.Printf("Could not log the user in: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Could not log the user in: %v", err))
	}
	log.Print(user)
	return c.JSON(http.StatusOK, user)
}
