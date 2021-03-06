package store

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	spotifyR "github.com/avidreder/monmach-api/resources/spotify"

	"github.com/fatih/structs"
	"github.com/zmb3/spotify"
	"gopkg.in/mgo.v2/bson"
)

// Store is the interface for data storage, such as a db
type Store interface {
	Connect() error
	AdminGetUser(interface{}, string, interface{}) error
	GetAll(bson.ObjectId, string, interface{}) error
	GetByKey(bson.ObjectId, string, interface{}, string, interface{}) error
	GetManyByKey(bson.ObjectId, string, interface{}, string, interface{}) error
	UpdateByKey(bson.ObjectId, string, map[string]interface{}, string, interface{}) error
	DeleteByKey(bson.ObjectId, string, string, interface{}) error
	Create(string, map[string]interface{}) error
	CountByQuery(bson.ObjectId, string, string, interface{}) (int, error)
}

// ValidateRequired checks all fields are present for CRUD operations
func ValidateRequired(schema interface{}, values map[string]interface{}) (map[string]interface{}, error) {
	newValues := map[string]interface{}{}
	structMap := structs.Map(schema)
	log.Printf("%+v", values)
	for k, v := range structMap {
		_, ok := values[k]
		if k != "ID" {
			if !ok {
				return nil, fmt.Errorf("Required field %s was not present", k)
			}
			if reflect.TypeOf(values[k]).String() == reflect.TypeOf(v).String() {
				newValues[strings.ToLower(k)] = values[k]
			} else if reflect.TypeOf(v).String() == "int64" {
				number, err := strconv.ParseInt(values[k].(string), 10, 64)
				if err != nil {
					return nil, fmt.Errorf("Required field %s was not present: %v, %v", k, err, values[k])
				}
				newValues[strings.ToLower(k)] = number
			} else if reflect.TypeOf(v).String() == "[]string" {
				if reflect.TypeOf(values[k]).String() == "string" {
					var array []string
					if (values[k].(string)) != "" {
						err := json.Unmarshal([]byte(values[k].(string)), &array)
						if err == nil {
							newValues[strings.ToLower(k)] = array
						} else {
							return nil, fmt.Errorf("Required field %s was not present: %v, %v", k, err, values[k])
						}
					} else {
						newValues[strings.ToLower(k)] = array
					}
				} else {
					newValues[k] = make([]string, 2)
				}
			} else if reflect.TypeOf(v).String() == "[]float64" {
				if reflect.TypeOf(values[k]).String() == "string" {
					var array []float64
					if (values[k].(string)) != "" {
						err := json.Unmarshal([]byte(values[k].(string)), &array)
						if err == nil {
							newValues[strings.ToLower(k)] = array
						} else {
							return nil, fmt.Errorf("Required field %s was not present", k)
						}
					} else {
						newValues[strings.ToLower(k)] = array
					}
				} else {
					newValues[k] = make([]float64, 2)
				}
			} else if k == "SpotifyTrack" {
				if reflect.TypeOf(values[k]).String() == "string" {
					track := spotifyR.SpotifyTrack{}
					if (values[k].(string)) != "" {
						err := json.Unmarshal([]byte(values[k].(string)), &track)
						if err == nil {
							newValues[strings.ToLower(k)] = track
						} else {
							return nil, fmt.Errorf("Required field %s was not present", k)
						}
					} else {
						newValues[strings.ToLower(k)] = track
					}
				} else {
					newValues[k] = spotifyR.SpotifyTrack{}
				}
			} else if k == "Features" {
				if reflect.TypeOf(values[k]).String() == "string" {
					features := spotify.AudioFeatures{}
					if (values[k].(string)) != "" {
						err := json.Unmarshal([]byte(values[k].(string)), &features)
						if err == nil {
							newValues[strings.ToLower(k)] = features
						} else {
							return nil, fmt.Errorf("Required field %s was not present", k)
						}
					} else {
						newValues[strings.ToLower(k)] = features
					}
				} else {
					newValues[k] = spotify.AudioFeatures{}
				}
			}
		}
	}
	return newValues, nil
}

func ValidateInputs(schema interface{}, values map[string]interface{}) map[string]interface{} {
	newValues := map[string]interface{}{}
	structMap := structs.Map(schema)
	log.Print(structMap)
	for k, v := range values {
		_, ok := structMap[k]
		if ok && reflect.TypeOf(structMap[k]).String() == reflect.TypeOf(v).String() {
			newValues[strings.ToLower(k)] = v
		} else if ok && reflect.TypeOf(structMap[k]).String() == "[]string" && reflect.TypeOf(v).String() == "string" {
			var array []string
			err := json.Unmarshal([]byte(v.(string)), &array)
			if err == nil {
				newValues[strings.ToLower(k)] = array
			}
		} else if ok && reflect.TypeOf(structMap[k]).String() == "[]float64" && reflect.TypeOf(v).String() == "string" {
			var array []float64
			err := json.Unmarshal([]byte(v.(string)), &array)
			if err == nil {
				newValues[strings.ToLower(k)] = array
			}
		}
	}
	log.Print(newValues)
	return newValues
}
