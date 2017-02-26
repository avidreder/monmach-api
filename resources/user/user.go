package user

import (
	"time"

	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
)

// User is the struct representing a user
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	SpotifyID      string        `bson:"spotifyid" json:"UserID"`
	Name           string        `bson:"name"`
	Email          string        `bson:"email"`
	AvatarURL      string        `bson:"avatarurl"`
	Token          oauth2.Token  `bson:"token"`
	Created        time.Time     `bson:"created"`
	Updated        time.Time     `bson:"updated"`
	TrackBlacklist []string      `bson:"trackblacklist"`
}
