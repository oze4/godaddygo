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

// NewConfig creates a config using `http.DefaultClient` as our client
//  - key is the api key
//  - secret is the api secret
//  - env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
//  - version should be `v1` or `v2`, use APIVersion1 or APIVersion2
func NewConfig(key, secret, env, version string) *Config {
	return &Config{
		Client:  http.DefaultClient,
		key:     key,
		secret:  secret,
		version: version,
		env:     env,
	}
}

// SetClient allows you to specify your own http client
func (c *Config) SetClient(h *http.Client) {
	c.Client = h
}

func (c *Config) make(method string, path string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
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
