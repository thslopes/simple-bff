package apicall

type ApiCall struct {
	Url         string
	QueryParams []QueryParam
}

type QueryParam struct {
	Name  string
	Value string
	Type  string
}

type Caller struct {
	Getter Getter
}

func (c *Caller) Get(apiCall ApiCall, queryString, headers map[string]string) ([]byte, error) {
	url := apiCall.Url

	queryParams := ""
	for _, v := range apiCall.QueryParams {
		switch v.Type {
		case "constant":
			queryParams += v.Name + "=" + v.Value + "&"
		case "querystring":
			queryParams += v.Name + "=" + queryString[v.Value] + "&"
		case "header":
			queryParams += v.Name + "=" + headers[v.Value] + "&"
		}
	}
	if queryParams != "" {
		url += "?" + queryParams[:len(queryParams)-1]
	}

	return c.Getter.Get(url)
}
