package postgres

import (
	// "errors"
	// "fmt"
	"log"
	// "strings"

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

// // GetAllPlaylists grabs all data from a table
// func (s *Store) GetAllPlaylists(table string) ([]playlistR.Playlist, error) {
// 	var results []playlistR.Playlist
// 	var result playlistR.Playlist
// 	query := fmt.Sprintf("SELECT * FROM %s;", table)
// 	rows, err := s.db.Query(query)
// 	defer rows.Close()
// 	if err != nil {
// 		log.Printf("Error from GetAllPlaylists: %v", err)
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&result.ID, &result.Name, &result.UserID, &result.Tracks, &result.Created, &result.Updated)
// 		if err != nil {
// 			log.Printf("Error from GetAllPlaylists while reading row from table: %s", table)
// 		}
// 		results = append(results, result)
// 	}
// 	log.Print(results)
// 	return results, nil
// }

// Create inserts a row into a table
func (s *Store) Create(model interface{}) error {
	err := s.db.Insert(model)
	if err != nil {
		log.Printf("Error from Create: %v", err)
		return err
	}
	return nil
}

// // Update updates an existing row in a table
// func (s *Store) Update(table string, id int64, valueMap map[string]interface{}) (string, error) {
// 	var keys []string
// 	var values []interface{}
// 	var vars []string
// 	count := 0
// 	for k, v := range valueMap {
// 		count++
// 		keys = append(keys, k)
// 		values = append(values, v)
// 		vars = append(vars, fmt.Sprintf("$%v", count))
// 	}
// 	keys = append(keys, "updated")
// 	vars = append(vars, "current_timestamp")
// 	keyString := strings.Join(keys, ", ")
// 	query := fmt.Sprintf("UPDATE %s SET (%s) = (%s) WHERE id = %v;", table, keyString, strings.Join(vars, ", "), id)
// 	stmt, err := s.db.Prepare(query)
// 	defer stmt.Close()
// 	if err != nil {
// 		log.Printf("Error from Update: %v", err)
// 		return "", err
// 	}
// 	res, err := stmt.Exec(values...)
// 	if err != nil {
// 		log.Printf("Error from Update: %v", err)
// 		return "", err
// 	}
// 	log.Printf("%+v", res)
// 	id, err = res.RowsAffected()
// 	if err != nil {
// 		log.Printf("Error from Update: %v", err)
// 		return "", err
// 	}
// 	return string(id), nil
// }

// // Delete deletes data from a table
// func (s *Store) Delete(table string, id int64) error {
// 	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", table)
// 	stmt, err := s.db.Prepare(query)
// 	defer stmt.Close()
// 	if err != nil {
// 		return err
// 	}
// 	res, err := stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}
// 	rows, err := res.RowsAffected()
// 	if err != nil || rows != 1 {
// 		return errors.New("Error deleting data")
// 	}
// 	return nil
// }

// Get returns a postgres instance
func Get() (*Store, error) {
	return &dataStore, nil
}

// Set sets the store (mostly for testing)
func Set(s Store) {
	dataStore = s
}
