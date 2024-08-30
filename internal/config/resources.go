package config

import (
	"encoding/json"
	"log"
	"os"
)

type Resource struct {
	Url    string
	Method string
}

var Resources map[string]Resource

func LoadResources() {
	data, err := os.ReadFile("configs/resources.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &Resources)
	if err != nil {
		log.Fatal(err)
	}
}
