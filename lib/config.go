package lib

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Urls []string `json:"urls"`
}

func ReadConfig(path string) Config {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	var data Config
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Println("Error reading config,", err)
		os.Exit(2)
	}

	log.Println("Config: ", data)

	return data
}
