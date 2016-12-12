package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
var dataStore = Store{}

// Store implements store interface
type Store struct {
	mongoSession *mgo.Session
}

const dbURL = "mongodb://localhost:27017"
const db = "monmach"

func (s Store) Connect() error {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return err
	}
	s.mongoSession = session
	return nil
}

func (s Store) GetAll(collection string, model interface{}) error {
	err := s.mongoSession.DB(db).C(collection).Find(bson.M{}).All(model)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) GetByKey(collection string, model interface{}, key string, value interface{}) error {
	err := s.mongoSession.DB(db).C(collection).Find(bson.M{key: value}).One(model)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) UpdateByKey(collection string, updates map[string]interface{}, key string, value interface{}) error {

	err := s.mongoSession.DB(db).C(collection).Update(bson.M{key: value}, bson.M{"$set":updates})
	if err != nil {
		return err
	}
	return nil
}

func (s Store) DeleteByKey(collection string, key string, value interface{}) error {
	err := s.mongoSession.DB(db).C(collection).Remove(bson.M{key: value})
	if err != nil {
		return err
	}
	return nil
}

func (s Store) Create(collection string, model interface{}) error {
	err := s.mongoSession.DB(db).C(collection).Insert(model)
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
