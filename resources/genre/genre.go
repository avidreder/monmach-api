package genre

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	UserID         string
	QueueID        string
	Name           string
	Description    string
	SeedArtists    []string
	SeedTracks     []string
	SeedPlaylists  []string
	AvatarURL      string
	Created        time.Time
	Updated        time.Time
	TrackBlacklist []string
	TrackWhitelist []string
}
