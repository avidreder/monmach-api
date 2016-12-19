package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dataStore = &Store{}

// Store implements *Store interface
type Store struct {
	Session *mgo.Session
}

const dbURL = "mongodb://localhost:27017"
const db = "monmach"

func (s *Store) Connect() error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	err = session.Ping()
	if err != nil {
		return err
	}
	s.Session = session
	return nil
}

func getCollection(database *mgo.Database, collectionName string) *mgo.Collection {
	collection := database.C(collectionName)
	return collection
}

func (s *Store) GetAll(collection string, model interface{}) error {
	log.Printf("GetAll: collection: %s, model: %T", collection, model)
	session := s.Session.Copy()
	defer session.Close()
	database := session.DB(db)
	c := getCollection(database, collection)
	err := c.Find(bson.M{}).All(model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetByKey(collection string, model interface{}, key string, value interface{}) error {
	log.Printf("Get: collection: %s, model: %+v", collection, model)
	session := s.Session.Copy()
	defer session.Close()
	database := session.DB(db)
	c := getCollection(database, collection)
	err := c.Find(bson.M{key: value}).One(model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateByKey(collection string, updates map[string]interface{}, key string, value interface{}) error {
	log.Printf("Update: collection: %s, updates: %+v", collection, updates)
	session := s.Session.Copy()
	defer session.Close()
	database := session.DB(db)
	c := getCollection(database, collection)
	bsonUpdates := bson.M{}
	for k, v := range updates {
		bsonUpdates[k] = v
	}
	err := c.Update(bson.M{key: value}, bson.M{"$set": bsonUpdates})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteByKey(collection string, key string, value interface{}) error {
	log.Printf("Delete: collection: %s, id: %+v", collection, value)
	session := s.Session.Copy()
	defer session.Close()
	database := session.DB(db)
	c := getCollection(database, collection)
	err := c.Remove(bson.M{key: value})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Create(collection string, values map[string]interface{}) error {
	log.Printf("Create: collection: %s, values: %+v", collection, values)
	session := s.Session.Copy()
	defer session.Close()
	database := session.DB(db)
	c := getCollection(database, collection)
	bsonValues := bson.M{}
	for k, v := range values {
		bsonValues[k] = v
	}
	err := c.Insert(bsonValues)
	if err != nil {
		return err
	}
	return nil
}

// Get returns a mongodb instance
func Get() (*Store, error) {
	return dataStore, nil
}

// Set sets the store (mostly for testing)
func Set(s *Store) {
	dataStore = s
}
