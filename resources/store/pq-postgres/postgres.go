package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	playlistR "github.com/avidreder/monmach-api/resources/playlist"

	// Following the lib pq example
	_ "github.com/lib/pq"
)

// PGArgs are the info used for the connection
const PGArgs = "user=monmach dbname=monmach sslmode=disable"

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
func (s *Store) Get(table string, id int64, values ...interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1;", table)
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Printf("Error from Get: %v", err)
		return err
	}
	err = stmt.QueryRow(id).Scan(values...)
	if err != nil {
		log.Printf("Error from Get: %v", err)
		return err
	}
	return nil
}

// GetAllPlaylists grabs all data from a table
func (s *Store) GetAllPlaylists(table string) ([]playlistR.Playlist, error) {
	var results []playlistR.Playlist
	var result playlistR.Playlist
	query := fmt.Sprintf("SELECT * FROM %s;", table)
	rows, err := s.db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Printf("Error from GetAllPlaylists: %v", err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.UserID, &result.Tracks, &result.Created, &result.Updated)
		if err != nil {
			log.Printf("Error from GetAllPlaylists while reading row from table: %s", table)
		}
		results = append(results, result)
	}
	log.Print(results)
	return results, nil
}

// Create inserts a row into a table
func (s *Store) Create(table string, valueMap map[string]interface{}) (string, error) {
	var keys []string
	var values []interface{}
	var vars []string
	count := 0
	log.Print(valueMap)
	for k, v := range valueMap {
		count++
		keys = append(keys, k)
		values = append(values, v)
		vars = append(vars, fmt.Sprintf("$%v", count))
	}
	keys = append(keys, "created")
	keys = append(keys, "updated")
	vars = append(vars, "current_timestamp")
	vars = append(vars, "current_timestamp")
	keyString := strings.Join(keys, ", ")
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s) ;", table, keyString, strings.Join(vars, ", "))
	log.Print(query)
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		log.Printf("Error from Create: %v", err)
		return "", err
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		log.Printf("Error from Create: %v", err)
		return "", err
	}
	log.Printf("%+v", res)
	id, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error from Create: %v", err)
		return "", err
	}
	return string(id), nil
}

// Update updates an existing row in a table
func (s *Store) Update(table string, id int64, valueMap map[string]interface{}) (string, error) {
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
		return "", err
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		log.Printf("Error from Update: %v", err)
		return "", err
	}
	log.Printf("%+v", res)
	id, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error from Update: %v", err)
		return "", err
	}
	return string(id), nil
}

// Delete deletes data from a table
func (s *Store) Delete(table string, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", table)
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows != 1 {
		return errors.New("Error deleting data")
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
