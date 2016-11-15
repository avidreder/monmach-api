package user

import (
	"time"
)

// User is the struct representing a user
type User struct {
	ID                  int64
	Name                string `sql:",pk"`
	Email               string
	AvatarURL           string
	SpotifyToken        string `json:"AccessToken"`
	SpotifyRefreshToken string `json:"RefreshToken"`
	Created             time.Time
	Updated             time.Time
	TrackBlacklist      []int64 `pg:",array"`
}
