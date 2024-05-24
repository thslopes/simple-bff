package apicall

type ApiCall struct {
	Url string
}

type Caller struct {
	Getter Getter
}

func (c *Caller) Get(apiCall ApiCall) ([]byte, error) {
	return c.Getter.Get(apiCall.Url)
}
