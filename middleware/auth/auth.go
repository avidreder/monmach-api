package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/avidreder/monmach-api/resources/auth"
	"github.com/markbates/goth/providers/spotify"

	"github.com/labstack/echo"
)

// LoadSpotifyProvider places initialized provider in the context for later use
func LoadSpotifyProvider(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := os.Open("/srv/monmach-api/spotify.json") // For read access.
		if err != nil {
			log.Printf("Could not Initialize Spotify: %s", err)
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Could not Initialize Spotify: %s", err)
		}
		credentials := auth.SpotifyCredentials{}
		err = json.Unmarshal(contents, &credentials)
		if err != nil {
			log.Printf("Could not Initialize Spotify: %s", err)
		}
		provider := spotify.New(credentials.ClientKey, credentials.Secret, credentials.CallbackURL, auth.SpotifyScopes...)
		c.Set("spotifyProvider", provider)
		return h(c)
	}

}

// GetSpotifyProvider retieves provider from the context
func GetSpotifyProvider(c echo.Context) *spotify.Provider {
	return c.Get("spotifyProvider").(*spotify.Provider)
}
