package store

import (
	"encoding/json"
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
	Create(string, interface{}) error
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
		}
	}
	log.Print(newValues)
	return newValues
}
