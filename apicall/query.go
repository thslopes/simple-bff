package apicall

import "encoding/json"

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

func (c *Caller) Get(query string, queryString, headers map[string]string) (interface{}, error) {
	apiCall := Queries[query]
	url := Resources[apiCall.Resource]

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

	body, err := c.Getter.Get(url, queryParams, pathParams, headersParams)
	return parseResult(body, apiCall.Returns), err
}

func parseResult(data []byte, results []string) map[string]interface{} {
	res := map[string]interface{}{}

	var mapData map[string]interface{}
	_ = json.Unmarshal(data, &mapData)

	for _, v := range results {
		res[v] = mapData[v]
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
