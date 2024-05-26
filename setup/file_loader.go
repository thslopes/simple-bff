package setup

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thslopes/bff/apicall"
)

var (
	queriesPath   = "queries/query.json"
	resourcesPath = "resources/resources.json"
)

type LoadFileErr struct {
	Err string
}

func (e *LoadFileErr) Error() string {
	return fmt.Sprintf("File not found (%s)", e.Err)
}

func LoadQueries() (map[string]apicall.Query, error) {
	plan, err := os.ReadFile(queriesPath)
	if err != nil {
		return nil, &LoadFileErr{Err: err.Error()}
	}
	var data map[string]apicall.Query
	err = json.Unmarshal(plan, &data)
	if err != nil {
		return nil, &LoadFileErr{Err: err.Error()}
	}
	return data, nil
}

func LoadResources() (map[string]apicall.Resource, error) {
	plan, err := os.ReadFile(resourcesPath)
	if err != nil {
		return nil, &LoadFileErr{Err: err.Error()}
	}
	var data map[string]apicall.Resource
	err = json.Unmarshal(plan, &data)
	if err != nil {
		return nil, &LoadFileErr{Err: err.Error()}
	}
	return data, nil
}
