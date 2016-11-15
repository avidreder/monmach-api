package track

import (
	"time"
)

// Track is the struct for a user's playlist
type Track struct {
	ID        int64
	Name      string
	Artists   []string `pg:",array"`
	ImageURL  string
	SpotifyID string
	Created   time.Time
	Updated   time.Time
	Features  []float64 `pg:",array"`
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
