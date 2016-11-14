package user

import (
	"net/url"
	"time"
)

// User is the struct representing a user
type User struct {
	ID             int64
	Name           string
	Email          string
	AvatarURL      url.URL
	Created        time.Time
	Updated        time.Time
	TrackBlacklist []int64
	TrackWhitelist []int64
	ListenedTracks []int64
}
