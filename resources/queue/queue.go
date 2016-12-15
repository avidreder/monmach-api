package queue

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Queue is the struct for a user's custom genre
type Queue struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	UserID         string        `bson:"userid,omitempty"`
	Name           string        `bson:"name,omitempty"`
	MaxSize        int64         `bson:"maxsize,omitempty"`
	TrackQueue     []string      `bson:"trackqueue,omitempty"`
	SeedArtists    []string      `bson:"seedartists,omitempty"`
	SeedTracks     []string      `bson:"seedtracks,omitempty"`
	ListenedTracks []string      `bson:"listenedtracks,omitempty"`
	Created        time.Time     `bson:"created,omitempty"`
	Updated        time.Time     `bson:"updated,omitempty"`
}
