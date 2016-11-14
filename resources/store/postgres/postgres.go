package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

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

type testy struct {
	ID   int64
	Name string
	Age  int64
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
		return err
	}
	err = stmt.QueryRow(id).Scan(values...)
	if err != nil {
		return err
	}

	log.Print(values...)
	return nil
}

// GetAllPlaylists grabs all data from a table
func (s *Store) GetAllPlaylists(table string) ([]testy, error) {
	var results []testy
	var result testy
	query := fmt.Sprintf("SELECT * FROM %s;", table)
	rows, err := s.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Age)
		if err != nil {
			log.Printf("Error reading row from table: %s", table)
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
	for k, v := range valueMap {
		count++
		keys = append(keys, k)
		values = append(values, v)
		vars = append(vars, fmt.Sprintf("$%v", count))
	}
	keyString := strings.Join(keys, ", ")
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s) RETURNING id;", table, keyString, strings.Join(vars, ", "))
	log.Print(query)
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		return "", err
	}
	log.Printf("%+v", res)
	id, err := res.RowsAffected()
	if err != nil {
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
	keyString := strings.Join(keys, ", ")
	query := fmt.Sprintf("UPDATE %s SET (%s) = (%s) WHERE id = %v;", table, keyString, strings.Join(vars, ", "), id)
	log.Print(query)
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(values...)
	if err != nil {
		return "", err
	}
	log.Printf("%+v", res)
	id, err = res.RowsAffected()
	if err != nil {
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
