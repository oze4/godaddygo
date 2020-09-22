package godaddygo

import (
	"fmt"
	"io"
	"net/http"
)

var (
	// ErrorWrongStatusCode is the error generated when an incorrect http status code is received
	ErrorWrongStatusCode = fmt.Errorf("ErrorWrongStatusCode")
)

// Config holds connection options
type Config struct {
    Client     *http.Client
	key        string
	secret     string
	baseURL    string
	env        string
	version    string
	path       string
	domainName string
}

func (c *Config) makeURLBase(xtra string) string {
	return c.baseURL + ""
}

// NewConfig creates a config using `http.DefaultClient` as our client
func NewConfig(key, secret, env, version string) *Config {
	return &Config{
		key:     key,
		secret:  secret,
		version: version,
		env:     env,
		Client:  http.DefaultClient,
	}
}

// SetClient allows you to specify your own http client
func (c *Config) SetClient(httpclient *http.Client) {
	c.Client = httpclient
}

func (c *Config) makeRequest(method string, path string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.baseURL+"/"+c.version+path, body)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %w", err)
	}
	req.Header.Set("Authorization", "sso-key "+c.key+":"+c.secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	if resp.StatusCode != expectedStatus {
		return resp.Body, fmt.Errorf("%w :expectedStatus %d, got %d", ErrorWrongStatusCode, expectedStatus, resp.StatusCode)
	}
	return resp.Body, nil
}
