package auth

type SpotifyCredentials struct {
	ClientKey   string `json:"clientKey"`
	Secret      string `json:"secret"`
	CallbackURL string `json:"callbackURL"`
}

var SpotifyScopes = []string{"user-read-email", "playlist-read-private", "playlist-modify-public", "playlist-modify-private"}

type SpotifyUser struct {
	Email        string
	Name         string
	UserID       string
	AvatarURL    string
	AccessToken  string
	RefreshToken string
}
