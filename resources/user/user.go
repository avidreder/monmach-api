package user

import (
	"time"
)

// User is the struct representing a user
type User struct {
	ID                  int64
	Name                string
	Email               string
	AvatarURL           string
	SpotifyToken        string
	SpotifyRefreshToken string
	Created             time.Time
	Updated             time.Time
	TrackBlacklist      []int64 `pg:",array"`
}
