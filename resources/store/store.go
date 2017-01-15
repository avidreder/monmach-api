package store

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// Store is the interface for data storage, such as a db
type Store interface {
	Connect() error
	GetAll(string, interface{}) error
	GetByKey(string, interface{}, string, interface{}) error
	UpdateByKey(string, map[string]interface{}, string, interface{}) error
	DeleteByKey(string, string, interface{}) error
	Create(string, map[string]interface{}) error
	CountByQuery(string, string, interface{}) (int, error)
}

// ValidateRequired checks all fields are present for CRUD operations
func ValidateRequired(schema interface{}, values map[string]interface{}) (map[string]interface{}, error) {
	newValues := map[string]interface{}{}
	structMap := structs.Map(schema)
	for k, v := range structMap {
		_, ok := values[k]
		if k != "ID" {
			if !ok {
				return nil, fmt.Errorf("Required field %s was not present", k)
			}
			if reflect.TypeOf(values[k]).String() == reflect.TypeOf(v).String() {
				newValues[k] = v
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
					log.Printf("Making new slice for: %s", k)
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
