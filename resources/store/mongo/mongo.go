package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dataStore = Store{}

// Store implements store interface
type Store struct {
	db *mgo.Database
}

const dbURL = "mongodb://localhost:27017"
const db = "monmach"

func (s Store) Connect() error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	err = session.Ping()
	if err != nil {
		return err
	}
	s.db = session.DB(db)
	return nil
}

func getCollection(database *mgo.Database, collectionName string) *mgo.Collection {
	collection := database.C(collectionName)
	log.Printf("collection: %+v", collection)
	return collection
}

func (s Store) GetAll(collection string, model interface{}) error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	defer session.Close()
	s.db = session.DB(db)
	c := getCollection(s.db, collection)
	err = c.Find(bson.M{}).All(model)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) GetByKey(collection string, model interface{}, key string, value interface{}) error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	defer session.Close()
	s.db = session.DB(db)
	c := getCollection(s.db, collection)
	err = c.Find(bson.M{key: value}).One(model)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) UpdateByKey(collection string, updates map[string]interface{}, key string, value interface{}) error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	defer session.Close()
	s.db = session.DB(db)
	c := getCollection(s.db, collection)
	err = c.Update(bson.M{key: value}, bson.M{"$set": updates})
	if err != nil {
		return err
	}
	return nil
}

func (s Store) DeleteByKey(collection string, key string, value interface{}) error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	defer session.Close()
	s.db = session.DB(db)
	c := getCollection(s.db, collection)
	err = c.Remove(bson.M{key: value})
	if err != nil {
		return err
	}
	return nil
}

func (s Store) Create(collection string, model interface{}) error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	defer session.Close()
	s.db = session.DB(db)
	c := getCollection(s.db, collection)
	err = c.Insert(model)
	if err != nil {
		return err
	}
	return nil
}

// Get returns a mongodb instance
func Get() (Store, error) {
	return dataStore, nil
}

// Set sets the store (mostly for testing)
func Set(s Store) {
	dataStore = s
}
