package apicall

import (
	"github.com/gofiber/fiber/v3/client"
)

type GetterErr struct {
	Err string
}

func (e *GetterErr) Error() string {
	return e.Err
}

type Getter interface {
	Do(resource Resource, qs, pathParams, headers map[string]string, body interface{}) ([]byte, error)
}

type FakeGetter struct {
	Error                   bool
	Url                     string
	Qs, PathParams, Headers map[string]string
	Body                    interface{}
}

func (f *FakeGetter) Do(resource Resource, qs, pathParams, headers map[string]string, body interface{}) ([]byte, error) {
	if f.Error {
		return nil, &GetterErr{Err: resource.Url}
	}
	f.Url = resource.Url
	f.Qs = qs
	f.PathParams = pathParams
	f.Headers = headers
	f.Body = body
	return []byte("{}"), nil
}

type httpGetter struct {
	Client *client.Client
}

func NewHttpGetter() Getter {
	return &httpGetter{
		Client: client.New(),
	}
}

func (h *httpGetter) Do(resource Resource, qs, pathParams, headers map[string]string, body interface{}) ([]byte, error) {
	resp, err := h.Client.Custom(resource.Url, resource.Method,
		client.Config{
			Param:     qs,
			PathParam: pathParams,
			Header:    headers,
			Body:      body,
		},
	)
	if err != nil {
		return nil, err
	}

	return resp.RawResponse.Body(), nil
}
