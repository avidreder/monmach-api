package queue

import (
	"time"

	"github.com/avidreder/monmach-api/resources/track"

	"gopkg.in/mgo.v2/bson"
)

// Queue is the struct for a user's custom genre
type Queue struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	UserID         bson.ObjectId `bson:"userid,omitempty"`
	Name           string        `bson:"name,omitempty"`
	MaxSize        int64         `bson:"maxsize,omitempty"`
	TrackQueue     []track.Track `bson:"trackqueue,omitempty"`
	SeedArtists    []string      `bson:"seedartists,omitempty"`
	SeedTracks     []string      `bson:"seedtracks,omitempty"`
	ListenedTracks []string      `bson:"listenedtracks,omitempty"`
	Created        time.Time     `bson:"created,omitempty"`
	Updated        time.Time     `bson:"updated,omitempty"`
}
