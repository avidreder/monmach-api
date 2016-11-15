package auth

import (
	authmw "github.com/avidreder/monmach-api/middleware/auth"

	"github.com/labstack/echo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

// StartAuth begins authorization
func StartAuth(c echo.Context) error {
	provider := authmw.GetSpotifyProvider(c)
	// c.SetParamNames("provider")
	// c.SetParamValues("spotify")
	q := c.Request().URL.Query()
	q.Add("provider", "spotify")
	c.Request().URL.RawQuery = q.Encode()
	goth.UseProviders(provider)
	gothic.BeginAuthHandler(c.Response().Writer(), c.Request())
	return nil
}
