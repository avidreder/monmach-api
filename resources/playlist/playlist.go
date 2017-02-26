package playlist

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Playlist is the struct for a user's playlist
type Playlist struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	OwnerID   string        `bson:"ownerid"`
	Name      string
	UserID    string
	SpotifyID string
	Tracks    []string
	Created   time.Time
	Updated   time.Time
}
