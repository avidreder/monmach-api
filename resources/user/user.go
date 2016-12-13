package user

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User is the struct representing a user
type User struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	SpotifyID           string        `json:"UserID"`
	Name                string
	Email               string
	AvatarURL           string
	SpotifyToken        string `json:"AccessToken"`
	SpotifyRefreshToken string `json:"RefreshToken"`
	Created             time.Time
	Updated             time.Time
	TrackBlacklist      []string
}
