package godaddygo

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// NewConfig creates a default config, allowing you to choose Production or
// Development GoDaddy API via `APIProdEnv` or `APIDevEnv` constants
func NewConfig(key, secret, env string) *Config {
	return newDefaultConfig(key, secret, env)
}

func newDefaultConfig(key, secret, env string) *Config {
	return &Config{
		client: http.DefaultClient,
		key:    key,
		secret: secret,
		env:    env,
	}
}

// Config holds connection options
type Config struct {
	client     *http.Client
	key        string // key is the api key
	secret     string // secret is the api secret
	baseURL    string // we take care of this
	env        string // env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
	version    string // version should be `v1` or `v2`, use APIVersion1 or APIVersion2
	domainName string // dns zone to target
}

// WithClient attaches a custom `*http.Client`
func (c *Config) WithClient(client *http.Client) {
	c.client = client
}

// makeDo makes an http.Request and sends it
func (c *Config) makeDo(ctx context.Context, method string, path string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.baseURL+"/"+c.version+path, body)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %w", err)
	}

	req.WithContext(ctx)

	req.Header.Set("Authorization", "sso-key "+c.key+":"+c.secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	if resp.StatusCode != expectedStatus {
		return resp.Body, fmt.Errorf("%w :expectedStatus %d, got %d", ErrorWrongStatusCode, expectedStatus, resp.StatusCode)
	}

	return resp.Body, nil
}
