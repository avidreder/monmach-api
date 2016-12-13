package queue

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Queue is the struct for a user's custom genre
type Queue struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	UserID         string
	Name           string
	MaxSize        int64
	TrackQueue     []string
	SeedArtists    []string
	SeedTracks     []string
	ListenedTracks []string
	Created        time.Time
	Updated        time.Time
}
