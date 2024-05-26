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
	Get(url string, qs, pathParams, headers map[string]string) ([]byte, error)
}

type FakeGetter struct {
	Error                   bool
	Url                     string
	Qs, PathParams, Headers map[string]string
}

func (f *FakeGetter) Get(url string, qs, pathParams, headers map[string]string) ([]byte, error) {
	if f.Error {
		return nil, &GetterErr{Err: url}
	}
	f.Url = url
	f.Qs = qs
	f.PathParams = pathParams
	f.Headers = headers
	return []byte("fake"), nil
}

type httpGetter struct {
	Client *client.Client
}

func NewHttpGetter() Getter {
	return &httpGetter{
		Client: client.New(),
	}
}

func (h *httpGetter) Get(url string, qs, pathParams, headers map[string]string) ([]byte, error) {
	resp, err := h.Client.Get(url, client.Config{
		Param:     qs,
		PathParam: pathParams,
		Header:    headers,
	})
	if err != nil {
		return nil, err
	}

	return resp.RawResponse.Body(), nil
}
