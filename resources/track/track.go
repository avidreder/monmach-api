package track

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Track is the struct for a user's playlist
type Track struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Artists   []string
	ImageURL  string
	SpotifyID string
	Genres    []string
	Playlists []string
	Rating    int64
	Created   time.Time
	Updated   time.Time
	Features  []float64
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

// Artist is a Spotify Artist profile
// type Artist struct {
// 	ID   string
// 	Name string
// }
