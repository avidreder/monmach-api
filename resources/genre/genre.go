package genre

import (
	"time"

	trackR "github.com/avidreder/monmach-api/resources/track"

	"gopkg.in/mgo.v2/bson"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             bson.ObjectId  `bson:"_id,omitempty"`
	UserID         string         `bson:"userid"`
	QueueID        string         `bson:"queueid"`
	Name           string         `bson:"name"`
	Description    string         `bson:"description"`
	SeedArtists    []string       `bson:"seedartists"`
	SeedTracks     []trackR.Track `bson:"seedtracks"`
	SeedPlaylists  []string       `bson:"seedplaylists"`
	AvatarURL      string         `bson:"avatarurl"`
	Created        time.Time      `bson:"created"`
	Updated        time.Time      `bson:"updated"`
	ListenedTracks []trackR.Track `bson:"listenedtracks"`
}
