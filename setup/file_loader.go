package setup

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thslopes/bff/apicall"
)

type LoadFileErr struct {
	Err string
}

func (e *LoadFileErr) Error() string {
	return fmt.Sprintf("File not found (%s)", e.Err)
}

func LoadQueries(filePath string) (map[string]apicall.Query, error) {
	plan, err := os.ReadFile(filePath)
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
