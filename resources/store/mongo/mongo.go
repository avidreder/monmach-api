package mongo

import (
	"gopkg.in/mgo.v2"
)

const dbURL = "mongodb://localhost:27017"
const db = "showhawk"
const collection = "shows"

func Connect() (*mgo.Collection, error) {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		return nil, err
	}
	col := session.DB(db).C(collection)
	return col, nil
}
