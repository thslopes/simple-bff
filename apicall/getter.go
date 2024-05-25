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
	Get(url string, qs map[string]string) ([]byte, error)
}

type FakeGetter struct {
	Error bool
}

func (f *FakeGetter) Get(url string, qs map[string]string) ([]byte, error) {
	if f.Error {
		return nil, &GetterErr{Err: url}
	}
	for k, v := range qs {
		url += k + v
	}
	return []byte(url), nil
}

type httpGetter struct {
	Client *client.Client
}

func NewHttpGetter() Getter {
	return &httpGetter{
		Client: client.New(),
	}
}

func (h *httpGetter) Get(url string, qs map[string]string) ([]byte, error) {
	resp, err := h.Client.Get(url, client.Config{
		Param: qs,
	})
	if err != nil {
		return nil, err
	}

	return resp.RawResponse.Body(), nil
}
