package playlist

import (
	"time"
)

// Playlist is the struct for a user's playlist
type Playlist struct {
	ID        int64
	Name      string
	UserID    int64
	SpotifyID string
	Tracks    []int64 `pg:",array"`
	Created   time.Time
	Updated   time.Time
}
