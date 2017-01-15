package queue

import (
	"time"

	"github.com/avidreder/monmach-api/resources/track"

	"gopkg.in/mgo.v2/bson"
)

// Queue is the struct for a user's custom genre
type Queue struct {
	ID             bson.ObjectId `bson:"_id"`
	UserID         bson.ObjectId `bson:"userid"`
	Name           string        `bson:"name"`
	MaxSize        int64         `bson:"maxsize"`
	TrackQueue     []track.Track `bson:"trackqueue"`
	SeedArtists    []string      `bson:"seedartists"`
	SeedTracks     []string      `bson:"seedtracks"`
	ListenedTracks []string      `bson:"listenedtracks"`
	Created        time.Time     `bson:"created"`
	Updated        time.Time     `bson:"updated"`
}
