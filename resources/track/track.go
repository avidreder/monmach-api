package track

import (
	"time"

	spotifyR "github.com/avidreder/monmach-api/resources/spotify"
	"github.com/avidreder/monmach-api/resources/store/mongo"

	"github.com/zmb3/spotify"
	"gopkg.in/mgo.v2/bson"
)

// Track is the struct for a user's playlist
type Track struct {
	ID           bson.ObjectId         `bson:"_id,omitempty"`
	SpotifyTrack spotifyR.SpotifyTrack `bson:"spotifytrack"`
	SpotifyID    string                `bson:"spotifyid"`
	OwnerID      bson.ObjectId         `bson:"ownerid"`
	Genres       []string              `bson:"genres"`
	CustomGenres []string              `bson:"customgenres"`
	Playlists    []string              `bson:"playlists"`
	Rating       int64                 `bson:"rating"`
	Created      time.Time             `bson:"created"`
	Updated      time.Time             `bson:"updated"`
	Features     spotify.AudioFeatures `bson:"features"`
}

// AlreadyProcessed checks if a track has already been added
func AlreadyProcessed(trackID string) (bool, error) {
	store, err := mongo.Get()
	if err != nil {
		return false, err
	}
	count, err := store.CountByQuery("tracks", "SpotifyTrack.SpotifyID", trackID)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// SpotifyFeatures is the response from Spotify's Audio Features API
// type SpotifyFeatures struct {
// 	Danceability     float64
// 	Energy           float64
// 	Key              int64
// 	Loudness         float64
// 	Mode             int64
// 	Speechiness      float64
// 	Acousticness     float64
// 	Instrumentalness float64
// 	Liveness         float64
// 	Valence          float64
// 	Tempo            float64
// 	Duration         int64
// 	TimeSignature    int64
// }
