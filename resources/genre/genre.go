package genre

import (
	"time"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             int64
	UserID         int64
	QueueID        int64
	Name           string
	Description    string
	SeedArtists    []int64 `pg:",array"`
	SeedTracks     []int64 `pg:",array"`
	AvatarURL      string
	Created        time.Time
	Updated        time.Time
	TrackBlacklist []int64 `pg:",array"`
	TrackWhitelist []int64 `pg:",array"`
}
