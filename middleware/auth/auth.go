package auth

import (
	"github.com/avidreder/monmach-api/resources/spotify"

	"github.com/labstack/echo"
	spotifyP "github.com/markbates/goth/providers/spotify"
)

// LoadSpotifyProvider places initialized provider in the context for later use
func LoadSpotifyProvider(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("spotifyProvider", spotify.SpotifyProvider)
		return h(c)
	}
}

// GetSpotifyProvider retieves provider from the context
func GetSpotifyProvider(c echo.Context) *spotifyP.Provider {
	return c.Get("spotifyProvider").(*spotifyP.Provider)
}
