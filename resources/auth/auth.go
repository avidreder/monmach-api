package auth

type SpotifyCredentials struct {
	ClientKey   string
	Secret      string
	CallbackURL string
}

var SpotifyScopes = []string{"user-read-email", "playlist-read-private", "playlist-modify-public", "playlist-modify-private"}
