package postgres

import (
	"database/sql"

	// Following the lib pq example
	_ "github.com/lib/pq"
)

// PGArgs are the info used for the connection
const PGArgs = "user=andrew dbname=monmach"

var dataStore = Store{}

// Store implements store interface
type Store struct {
	db *sql.DB
}

// Connect creates a connection to the postgres db
func (s *Store) Connect() error {
	db, err := sql.Open("postgres", PGArgs)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}

// Get returns a postgres instance
func Get() (*Store, error) {
	return &dataStore, nil
}
