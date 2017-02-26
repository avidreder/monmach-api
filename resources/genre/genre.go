package genre

import (
	"time"

	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	trackR "github.com/avidreder/monmach-api/resources/track"

	"gopkg.in/mgo.v2/bson"
)

// Genre is the struct for a user's custom genre
type Genre struct {
	ID             bson.ObjectId            `bson:"_id,omitempty"`
	UserID         bson.ObjectId            `bson:"userid"`
	OwnerID        bson.ObjectId            `bson:"ownerid"`
	QueueID        bson.ObjectId            `bson:"queueid"`
	Name           string                   `bson:"name"`
	Description    string                   `bson:"description"`
	SeedArtists    []spotifyR.SpotifyArtist `bson:"seedartists"`
	SeedTracks     []trackR.Track           `bson:"seedtracks"`
	SeedGenres     []string                 `bson:"seedgenres"`
	AvatarURL      string                   `bson:"avatarurl"`
	Created        time.Time                `bson:"created"`
	Updated        time.Time                `bson:"updated"`
	ListenedTracks []trackR.Track           `bson:"listenedtracks"`
}
