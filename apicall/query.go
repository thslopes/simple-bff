package apicall

import (
	"encoding/json"
	"strings"
)

type Resource struct {
	Url    string
	Method string
}

type QueryErr struct {
	Err error
}

func (e *QueryErr) Error() string {
	return e.Err.Error()
}

type Query struct {
	Resource    string
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

type Caller struct {
	Getter Getter
}

func (c *Caller) Do(query string, queryString, headers map[string]string, body interface{}) (interface{}, error) {
	apiCall := Queries[query]
	resource := Resources[apiCall.Resource]

	queryParams := map[string]string{}
	for _, v := range apiCall.QueryParams {
		queryParams[v.Name] = getParamValue(v, queryString, headers)
	}

	pathParams := map[string]string{}
	for _, v := range apiCall.PathParams {
		pathParams[v.Name] = getParamValue(v, queryString, headers)
	}

	headersParams := map[string]string{}
	for _, v := range apiCall.Headers {
		headersParams[v.Name] = getParamValue(v, queryString, headers)
	}

	respBody, err := c.Getter.Do(resource, queryParams, pathParams, headersParams, body)

	if err != nil {
		return nil, err
	}

	var mapData interface{}
	err = json.Unmarshal(respBody, &mapData)

	if err != nil {
		return nil, &QueryErr{Err: err}
	}

	return parseResult(mapData, apiCall.Returns), err
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

func getParamValue(v Param, queryString map[string]string, headers map[string]string) string {
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
