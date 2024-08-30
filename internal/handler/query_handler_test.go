package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestQueryHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/query/:query", QueryHandler)

	tests := []struct {
		name           string
		query          string
		queryParams    map[string]string
		headers        map[string]string
		requestBody    interface{}
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:           "success",
			query:          "example",
			queryParams:    map[string]string{"param1": "value1", "param2": "value2"},
			headers:        map[string]string{"Content-Type": "application/json"},
			requestBody:    map[string]interface{}{"key": "value"},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]interface{}{"result": "success"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodGet, "/query/"+tt.query, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			for key, value := range tt.queryParams {
				req.URL.Query().Add(key, value)
			}

			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var responseBody interface{}
			err = json.NewDecoder(resp.Body).Decode(&responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}
