package genre

import (
	"net/url"
	"time"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             int64
	UserID         int64
	Name           string
	Description    string
	SeedArtists    []int64
	SeedTracks     []int64
	AvatarURL      url.URL
	Created        time.Time
	Updated        time.Time
	TrackBlacklist []int64
	TrackWhitelist []int64
}
