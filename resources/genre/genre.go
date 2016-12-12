package genre

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	UserID         int64
	QueueID        int64
	Name           string
	Description    string
	SeedArtists    []int64 `pg:",array"`
	SeedTracks     []int64 `pg:",array"`
	SeedPlaylists  []int64 `pg:",array"`
	AvatarURL      string
	Created        time.Time
	Updated        time.Time
	TrackBlacklist []int64 `pg:",array"`
	TrackWhitelist []int64 `pg:",array"`
}
