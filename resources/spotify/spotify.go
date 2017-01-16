package spotify

import "github.com/zmb3/spotify"

type FeaturedPlaylist struct {
	Name       string
	OwnerID    spotify.ID
	PlaylistID spotify.ID
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
