package apicall

import (
	"fmt"

	"github.com/gofiber/fiber/v3/client"
)

type GetterErr struct {
	Err string
}

func (e *GetterErr) Error() string {
	return e.Err
}

type Getter interface {
	Get(string) ([]byte, error)
}

type FakeGetter struct {
	Error bool
}

func (f *FakeGetter) Get(url string) ([]byte, error) {
	if f.Error {
		return nil, &GetterErr{Err: url}
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

func (h *httpGetter) Get(url string) ([]byte, error) {
	fmt.Println(url)
	resp, err := h.Client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.RawResponse.Body(), nil
}
