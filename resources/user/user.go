package user

import (
	"time"

	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
)

// User is the struct representing a user
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	SpotifyID      string        `bson:"SpotifyID" json:"UserID"`
	Name           string        `bson:"Name"`
	Email          string        `bson:"Email"`
	AvatarURL      string        `bson:"AvatarURL"`
	Token          oauth2.Token  `bson:"Token"`
	Created        time.Time     `bson:"Created"`
	Updated        time.Time     `bson:"Updated"`
	TrackBlacklist []string      `bson:"TrackBlacklist"`
}
