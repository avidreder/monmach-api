package mongo

import (
	"log"
	"encoding/json"

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

func (s Store) GetAll(model interface{}, collection string) error {
	data, err := s.mongoSession.DB(db).Collection(collection).Find().All(model)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) GetByKey(model interface{}, key string, value interface{}) error {
	err := s.mongoSession.DB(db).Collection(collection).Find(bson.M{key: value}).One(model)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) UpdateByKey(model interface{}, key string, value interface{}) error {
	err := s.mongoSession.DB(db).Collection(collection).Update(bson.M{key: value}, bson.M{"$set":bson.M{model}})
	if err != nil {
		return err
	}
	return nil
}

func (s Store) DeleteByKey(model interface{}, key string, value interface{}) error {
	err := s.mongoSession.DB(db).Collection(collection).Remove{bson.M{key: value})
	if err != nil {
		return err
	}
	return nil
}

func (s Store) Create(model interface{}) error {
	err := s.mongoSession.DB(db).Collection(collection).Insert{model})
	if err != nil {
		return err
	}
	return nil
}
