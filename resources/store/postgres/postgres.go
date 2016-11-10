package postgres

import (
	"database/sql"
	// "database/sql/driver"
	// "github.com/lib/pq"
)

const PGTable = "postgres"
const PGArgs = "user=pqgotest dbname=pqgotest sslmode=verify-full"

type Store struct {
	db *sql.DB
}

func (s Store) Connect() error {
	db, err := sql.Open(PGTable, PGArgs)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}
