package apicall

type ApiCall struct {
	Url         string
	QueryParams []Param
	PathParams  []Param
}

type Param struct {
	Name  string
	Value string
	Type  string
}

type Caller struct {
	Getter Getter
}

func (c *Caller) Get(apiCall ApiCall, queryString, headers map[string]string) ([]byte, error) {
	url := apiCall.Url

	queryParams := map[string]string{}
	for _, v := range apiCall.QueryParams {
		queryParams[v.Name] = getParamValue(v, queryString, headers)
	}

	pathParams := map[string]string{}
	for _, v := range apiCall.PathParams {
		pathParams[v.Name] = getParamValue(v, queryString, headers)
	}

	return c.Getter.Get(url, queryParams, pathParams)
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
