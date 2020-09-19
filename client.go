package godaddygo

import (
	"fmt"
	"io"
	"net/http"
)

var (
    // ErrorWrongStatusCode is the error generated when an incorrect
    // http status code is received
	ErrorWrongStatusCode = fmt.Errorf("ErrorWrongStatusCode")
)

type client struct {
	key     string
	secret  string
	baseURL string
}

// New returns a new client
func NewClient(key string, secret string, baseURL string) *client {
	return &client{key, secret, baseURL}
}

func (c *client) make(method string, url string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.baseURL+url, body)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %w", err)
	}
	req.Header.Set("Authorization", "sso-key "+c.key+":"+c.secret)
	req.Header.Set("Content-Type", "application/json")

	httpclient := &http.Client{}
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	if resp.StatusCode != expectedStatus {
		return resp.Body, fmt.Errorf("%w :expectedStatus %d, got %d", ErrorWrongStatusCode, expectedStatus, resp.StatusCode)
	}
	return resp.Body, nil
}

func (c *client) Get(url string) (io.ReadCloser, error) {
	return c.make(http.MethodGet, url, nil, http.StatusOK)
}

func (c *client) Post(url string, body io.Reader) (io.ReadCloser, error) {
	return c.make(http.MethodPost, url, nil, http.StatusCreated)
}

func (c *client) Put(url string, body io.Reader) (io.ReadCloser, error) {
	return c.make(http.MethodPut, url, nil, http.StatusOK)
}

func (c *client) Delete(url string) error {
	result, err := c.make(http.MethodDelete, url, nil, http.StatusNoContent)
	defer result.Close()
	return err
}