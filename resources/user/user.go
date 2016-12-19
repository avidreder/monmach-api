package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User is the struct representing a user
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	SpotifyID      string        `bson:"SpotifyID"`
	Name           string        `bson:"Name"`
	Email          string        `bson:"Email"`
	AvatarURL      string        `bson:"AvatarURL"`
	AccessToken    string        `bson:"AccessToken"`
	RefreshToken   string        `bson:"RefreshToken"`
	Created        time.Time     `bson:"Created"`
	Updated        time.Time     `bson:"Updated"`
	TrackBlacklist []string      `bson:"TrackBlacklist"`
}
