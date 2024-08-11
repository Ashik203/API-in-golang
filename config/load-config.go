package config

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var filename = "/home/ashikurrahman/API/config.json"

func LoadConfig() (Config, error) {
	var config Config

	configFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error in Reading the config file")
		return config, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Println("Error in Parsing the config file")
	}

	return config, err
}
