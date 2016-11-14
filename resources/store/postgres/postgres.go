package postgres

import (
	// "errors"
	"fmt"
	"log"
	"strings"

	// playlistR "github.com/avidreder/monmach-api/resources/playlist"

	"gopkg.in/pg.v5"
)

var PGOpts = pg.Options{
	User:     "monmach",
	Database: "monmach",
}

var dataStore = Store{}

// Store implements store interface
type Store struct {
	db *pg.DB
}

// Connect creates a connection to the postgres db
func (s *Store) Connect() error {
	db := pg.Connect(&PGOpts)
	s.db = db
	return nil
}

// Get grabs data from a table
func (s *Store) Get(model interface{}) error {
	err := s.db.Select(model)
	if err != nil {
		log.Printf("Error from Get: %v", err)
		return err
	}
	return nil
}

// GetAll grabs all data from a table
func (s *Store) GetAll(model interface{}, tableName string) error {
	query := fmt.Sprintf(`SELECT * FROM %s`, tableName)
	_, err := s.db.Query(model, query)
	if err != nil {
		log.Printf("Error from GetAllPlaylists: %v", err)
		return err
	}
	return nil
}

// Create inserts a row into a table
func (s *Store) Create(model interface{}) error {
	err := s.db.Insert(model)
	if err != nil {
		log.Printf("Error from Create: %v", err)
		return err
	}
	return nil
}

// Update updates an existing row in a table
func (s *Store) Update(table string, id int64, valueMap map[string]interface{}) error {
	var keys []string
	var values []interface{}
	var vars []string
	count := 0
	for k, v := range valueMap {
		count++
		keys = append(keys, k)
		values = append(values, v)
		vars = append(vars, fmt.Sprintf("$%v", count))
	}
	keys = append(keys, "updated")
	vars = append(vars, "current_timestamp")
	keyString := strings.Join(keys, ", ")
	query := fmt.Sprintf("UPDATE %s SET (%s) = (%s) WHERE id = %v;", table, keyString, strings.Join(vars, ", "), id)
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Printf("Error from Update: %v", err)
		return err
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		log.Printf("Error from Update: %v", err)
		return err
	}
	return nil
}

// Delete deletes data from a table
func (s *Store) Delete(model interface{}) error {
	err := s.db.Delete(model)
	if err != nil {
		log.Printf("Error from Delete: %v", err)
		return err
	}
	return nil
}

// Get returns a postgres instance
func Get() (*Store, error) {
	return &dataStore, nil
}

// Set sets the store (mostly for testing)
func Set(s Store) {
	dataStore = s
}
