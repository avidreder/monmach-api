package queue

import (
	"time"
)

// Queue is the struct for a user's custom genre
type Queue struct {
	ID             int64
	UserID         int64 `db:"user_id"`
	Name           string
	MaxSize        int64
	TrackQueue     []int64 `pg:",array"`
	SeedArtists    []int64 `pg:",array"`
	SeedTracks     []int64 `pg:",array"`
	ListenedTracks []int64 `pg:",array"`
	Created        time.Time
	Updated        time.Time
}
