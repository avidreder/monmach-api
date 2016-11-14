package genre

import (
	"time"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             int64
	UserID         int64 `db:"user_id"`
	Name           string
	Description    string
	SeedArtists    []int64 `pg:",array" db:"seed_artists"`
	SeedTracks     []int64 `pg:",array" db:"seed_tracks"`
	AvatarURL      string  `db:"avatar_url"`
	Created        time.Time
	Updated        time.Time
	TrackBlacklist []int64 `pg:",array" db:"track_blacklist"`
	TrackWhitelist []int64 `pg:",array" db:"track_whitelist"`
}
