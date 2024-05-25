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
			got, err := LoadQueries(test.filePath)
			if diff := cmp.Diff(test.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}
