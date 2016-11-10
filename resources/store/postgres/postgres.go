package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// Following the lib pq example
	_ "github.com/lib/pq"
)

// PGArgs are the info used for the connection
const PGArgs = "user=andrew dbname=monmach sslmode=disable"

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

// Get grabs data from a table
func (s *Store) Get(table string, id string) (interface{}, error) {
	var result interface{}
	err := s.db.QueryRow("SELECT * FROM ? WHERE id=?", table, id).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAll grabs all data from a table
func (s *Store) GetAll(table string) ([]interface{}, error) {
	var results []interface{}
	var result interface{}
	rows, err := s.db.Query("SELECT * FROM ?", table)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			results = append(results, result)
			log.Printf("Error reading row from table: %s", table)
		}
	}
	return results, nil
}

// Create inserts a row into a table
func (s *Store) Create(table string, valueMap map[string]interface{}) (string, error) {
	var keys []string
	var values []interface{}
	var valueSlice []string
	for k, v := range valueMap {
		keys = append(keys, k)
		values = append(values, v)
		_, ok := v.(string)
		if ok {
			valueSlice = append(valueSlice, "'?'")
		} else {
			valueSlice = append(valueSlice, "?")
		}
	}
	keyString := strings.Join(keys, ", ")
	valueString := strings.Join(valueSlice, ",")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(?);", table, keyString)
	log.Print(valueString)
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(values)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return "", err
	}
	return string(id), nil
}

// Get returns a postgres instance
func Get() (*Store, error) {
	return &dataStore, nil
}

// Set sets the store (mostly for testing)
func Set(s Store) {
	dataStore = s
}
