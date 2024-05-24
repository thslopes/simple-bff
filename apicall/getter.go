package apicall

import (
	"io"
	"net/http"
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
	Client *http.Client
}

func NewHttpGetter() Getter {
	return &httpGetter{
		Client: &http.Client{},
	}
}

func (h *httpGetter) Get(url string) ([]byte, error) {
	resp, err := h.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
