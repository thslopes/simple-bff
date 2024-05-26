package setup

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thslopes/bff/apicall"
)

func TestLoadQueries(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		filePath string
		want     map[string]apicall.Query
		wantErr  error
	}{
		{
			name:     "success",
			filePath: "../testData/query.json",
			want: map[string]apicall.Query{
				"swapi-people": {
					Resource: "swapi-people",
					QueryParams: []apicall.Param{
						{Name: "format", Value: "json", Type: "constant"},
						{Name: "other", Value: "hello", Type: "querystring"},
					},
					PathParams: []apicall.Param{
						{Name: "personId", Value: "personId", Type: "querystring"},
					},
					Headers: []apicall.Param{
						{Name: "Authorization", Value: "Bearer 1234", Type: "constant"},
					},
					Returns: []string{"name"},
				},
			},
			wantErr: nil,
		},
		{
			name:     "file not found",
			filePath: "testData/tesFile.json",
			wantErr: &LoadFileErr{
				Err: "open testData/tesFile.json: no such file or directory",
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			queriesPath = test.filePath
			got, err := LoadQueries()
			if diff := cmp.Diff(test.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestLoadResources(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     map[string]string
		wantErr  error
	}{
		{
			name:     "success",
			filePath: "../testData/resource.json",
			want: map[string]string{
				"swapi-people": "https://swapi.dev/api/people/:personId/",
			},
			wantErr: nil,
		},
		{
			name:     "file not found",
			filePath: "testData/tesFile.json",
			wantErr: &LoadFileErr{
				Err: "open testData/tesFile.json: no such file or directory",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resourcesPath = tt.filePath
			got, err := LoadResources()
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}
