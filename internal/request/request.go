package request

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3/client"
	"github.com/thslopes/bff/internal/config"
)

type RequestErr struct {
	Err error
}

func (e *RequestErr) Error() string {
	return e.Err.Error()
}

type Request struct {
	Client               *client.Client
	Query                string
	QueryParams, Headers map[string]string
	Body                 interface{}
}

func NewRequest(query string) *Request {
	return &Request{
		Client:      client.New(),
		Query:       query,
		QueryParams: map[string]string{},
		Headers:     map[string]string{},
		Body:        nil,
	}
}

func (r *Request) Execute() (interface{}, error) {
	query := config.Queries[r.Query]
	resource := config.Resources[query.Resource]

	fmt.Println("Executing query", r.Query, "on resource", resource)

	queryParams := map[string]string{}
	for _, v := range query.QueryParams {
		queryParams[v.Name] = getParamValue(v, r.QueryParams, r.Headers)
	}

	pathParams := map[string]string{}
	for _, v := range query.PathParams {
		pathParams[v.Name] = getParamValue(v, r.QueryParams, r.Headers)
	}

	headersParams := map[string]string{}
	for _, v := range query.Headers {
		headersParams[v.Name] = getParamValue(v, r.QueryParams, r.Headers)
	}

	resp, err := r.Client.Custom(resource.Url, resource.Method,
		client.Config{
			Param:     queryParams,
			PathParam: pathParams,
			Header:    headersParams,
			Body:      r.Body,
		},
	)

	if err != nil {
		return nil, err
	}

	var mapData interface{}
	err = json.Unmarshal(resp.Body(), &mapData)

	if err != nil {
		return nil, &RequestErr{Err: err}
	}

	return parseResult(mapData, query.Returns), err
}

func parseResult(data interface{}, results []string) interface{} {
	res := map[string]interface{}{}

	for _, property := range results {
		props := strings.Split(property, ".")
		if props[0] == "[]" {
			resArray := []interface{}{}
			mapDataArray := data.([]interface{})
			for _, v := range mapDataArray {
				resArray = append(resArray, parseResult(v, []string{strings.Join(props[1:], ".")}))
			}
			return resArray
		}

		mapData := data.(map[string]interface{})
		if len(props) == 1 {
			if props[0] == "*" {
				return mapData
			}
			res[props[0]] = mapData[property]
		} else {
			res[props[0]] = parseResult(mapData[props[0]], []string{strings.Join(props[1:], ".")})
		}
	}

	return res
}

func getParamValue(v config.Param, queryString map[string]string, headers map[string]string) string {
	va := ""
	switch v.Type {
	case "constant":
		va = v.Value
	case "querystring":
		va = queryString[v.Value]
	case "header":
		va = headers[v.Value]
	}
	return va
}
