package xhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	netUrl "net/url"
	"time"

	xstring "github.com/dabao-zhao/helper/string"
)

type Http struct {
	Headers map[string]string
	Timeout time.Duration
}

type Options func(*Http)

func WithHeaders(headers map[string]string) Options {
	return func(http *Http) {
		http.Headers = headers
	}
}

func WithTimeout(timeout time.Duration) Options {
	return func(http *Http) {
		http.Timeout = timeout
	}
}

func NewHttp(opt ...Options) *Http {
	h := new(Http)
	for _, f := range opt {
		f(h)
	}
	return h
}

func (h *Http) Post(url string, data map[string]any) ([]byte, error) {
	param, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	if len(h.Headers) != 0 {
		for k, v := range h.Headers {
			req.Header.Set(k, v)
		}
	}

	client := new(http.Client)
	if h.Timeout != 0 {
		client.Timeout = h.Timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bt, nil
}

func (h *Http) Get(rawUrl string, data map[string]any) ([]byte, error) {
	url, err := netUrl.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	query := url.Query()
	for k, v := range data {
		query.Set(k, xstring.ToString(v))
	}
	url.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	if len(h.Headers) != 0 {
		for k, v := range h.Headers {
			req.Header.Set(k, v)
		}
	}

	client := new(http.Client)
	if h.Timeout != 0 {
		client.Timeout = h.Timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bt, nil
}
