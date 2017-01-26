package spotify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/avidreder/monmach-api/resources/auth"
	spotifyP "github.com/markbates/goth/providers/spotify"
	"github.com/zmb3/spotify"
)

type FeaturedPlaylist struct {
	Name       string
	OwnerID    spotify.ID
	PlaylistID spotify.ID
}

// SpotifyProvider stores an initialized provider
var SpotifyProvider *spotifyP.Provider

// InitializeSpotifyProvider places initialized provider in the context for later use
func InitializeSpotifyProvider() error {
	file, err := os.Open("/srv/monmach-api/spotify.json") // For read access.
	if err != nil {
		return err
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	credentials := auth.SpotifyCredentials{}
	err = json.Unmarshal(contents, &credentials)
	if err != nil {
		return err
	}
	log.Print(credentials)
	SpotifyProvider = spotifyP.New(credentials.ClientKey, credentials.Secret, credentials.CallbackURL, auth.SpotifyScopes...)
	return nil
}

var FeaturedPlaylists = []FeaturedPlaylist{
	FeaturedPlaylist{
		Name:       "Pitchfork Top Tracks",
		OwnerID:    "pitchforkmedia",
		PlaylistID: "5ItokEl1f0bbHeXWFiisrm",
	},
}

type SpotifyTrack struct {
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
		Name      string   `json:"name"`
		SpotifyID string   `json:"id"`
		Genres    []string `json:"genres"`
	} `json:"artists"`
}

type SpotifyResponse struct {
	Track SpotifyTrack `json:"track"`
}
