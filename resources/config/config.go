package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ServiceConfig struct {
	SpotifyCredentialsPath string
	MongoCredentialsPath   string
	ClientAddress          string
}

var CurrentConfig ServiceConfig

func init() {
	file, err := os.Open("/srv/monmach-api/localhost.json")
	if err != nil {
		log.Fatalf("Could not Initialize Serfice Config: %s", err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Could not Initialize Service Config: %s", err)
	}
	err = json.Unmarshal(contents, &CurrentConfig)
	if err != nil {
		log.Fatalf("Could not Initialize Service Config: %s", err)
	}
}

// LoadConfig pulls the config file from the file system
func GetConfig() (*ServiceConfig, error) {
	return &CurrentConfig, nil
}

// SetConfig set config (for testing)
func SetConfig(s ServiceConfig) {
	CurrentConfig = s
}
