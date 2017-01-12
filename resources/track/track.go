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
	SpotifyTrack spotifyR.SpotifyTrack `bson:"spotifytrack,omitempty"`
	SpotifyID    string                `bson:"spotifyid,omitempty"`
	Genres       []string              `bson:"genres,omitempty"`
	CustomGenres []string              `bson:"customgenres,omitempty"`
	Playlists    []string              `bson:"playlists,omitempty"`
	Rating       int64                 `bson:"rating,omitempty"`
	Created      time.Time             `bson:"created,omitempty"`
	Updated      time.Time             `bson:"updated,omitempty"`
	Features     spotify.AudioFeatures `bson:"features,omitempty"`
}

// AlreadyProcessed checks if a track has already been added
func AlreadyProcessed(trackID string) (bool, error) {
	store, err := mongo.Get()
	if err != nil {
		return false, err
	}
	count, err := store.CountByQuery("tracks", "SpotifyTrack.Track.SpotifyID", trackID)
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
