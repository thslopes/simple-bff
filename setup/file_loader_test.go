package setup

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thslopes/bff/apicall"
)

func TestLoadApiCallFromFile(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		filePath string
		want     map[string]apicall.ApiCall
		wantErr  error
	}{
		{
			name:     "success",
			filePath: "../testData/testFile.json",
			want: map[string]apicall.ApiCall{
				"swapi-people": {
					Url: "https://swapi.dev/api/people/1/",
					QueryParams: []apicall.QueryParam{{Name: "format", Value: "json", Type: "constant"}},
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
			got, err := LoadApiCallFromFile(test.filePath)
			if diff := cmp.Diff(test.wantErr, err); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("(-expected +actual):\n%s", diff)
			}
		})
	}
}
