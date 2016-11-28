package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	usermw "github.com/avidreder/monmach-api/middleware/user"
	authR "github.com/avidreder/monmach-api/resources/auth"

	"github.com/labstack/echo"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type SpotifyTrack struct {
	Track struct {
		Name      string `json:"name"`
		SpotifyID string `json:"id"`
		Album     struct {
			Images []struct {
				Height int64  `json:"height"`
				Width  int64  `json:"width"`
				URL    string `json:"url"`
			} `json:"images"`
		} `json:"album"`
		Artists []struct {
			Name      string `json:"name"`
			SpotifyID string `json:"id"`
		} `json:"artists"`
	} `json:"track"`
}

// LoadClient places initialized spotify client
func LoadClient(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := os.Open("/srv/monmach-api/spotify.json") // For read access.
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		credentials := authR.SpotifyCredentials{}
		err = json.Unmarshal(contents, &credentials)
		if err != nil {
			log.Printf("Could not Initialize Spotify Client: %s", err)
		}
		auth := spotify.NewAuthenticator(credentials.CallbackURL, spotify.ScopeUserReadPrivate)
		auth.SetAuthInfo(credentials.ClientKey, credentials.Secret)

		user := usermw.GetUser(c)
		token := &oauth2.Token{
			AccessToken: user.SpotifyToken,
		}
		client := auth.NewClient(token)
		c.Set("spotifyClient", &client)
		return h(c)
	}

}

// GetClient retieves provider from the context
func GetClient(c echo.Context) *spotify.Client {
	return c.Get("spotifyClient").(*spotify.Client)
}