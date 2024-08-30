package config

import (
	"encoding/json"
	"log"
	"os"
)

type Query struct {
	Resource    string `json:"resource"`
	QueryParams []Param
	PathParams  []Param
	Headers     []Param
	Returns     []string
}

type Param struct {
	Name  string
	Value string
	Type  string
}

var Queries map[string]Query

func LoadQueries() {
	data, err := os.ReadFile("configs/queries.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &Queries)
	if err != nil {
		log.Fatal(err)
	}
}
