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
}

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
			} else if ok && reflect.TypeOf(v).String() == "[]string" && reflect.TypeOf(values[k]).String() == "string" {
				var array []string
				err := json.Unmarshal([]byte(values[k].(string)), &array)
				if err == nil {
					newValues[strings.ToLower(k)] = array
				}
			} else if ok && reflect.TypeOf(structMap[k]).String() == "[]float64" && reflect.TypeOf(v).String() == "string" {
				var array []float64
				err := json.Unmarshal([]byte(v.(string)), &array)
				if err == nil {
					newValues[strings.ToLower(k)] = array
				} else {
					return nil, fmt.Errorf("Required field %s was not present", k)
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
		log.Printf("%v", k)
		log.Printf("%v,%T", v, v)
		_, ok := structMap[k]
		if ok {
			log.Print(reflect.TypeOf(structMap[k]).String())
		}
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
