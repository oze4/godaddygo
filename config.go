package godaddygo

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func newDefaultConfig(key, secret, env, version string) *Config {
	return &Config{http.DefaultClient, key, secret, "", env, version, ""}
}

// Config holds connection options
type Config struct {
	Client     *http.Client
	key        string // key is the api key
	secret     string // secret is the api secret
	baseURL    string // we take care of this
	env        string // env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
	version    string // version should be `v1` or `v2`, use APIVersion1 or APIVersion2
	domainName string // dns zone to target
}

// SetClient allows you to specify your own http client
func (c *Config) SetClient(h *http.Client) {
	c.Client = h
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

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	if resp.StatusCode != expectedStatus {
		return resp.Body, fmt.Errorf("%w :expectedStatus %d, got %d", ErrorWrongStatusCode, expectedStatus, resp.StatusCode)
	}

	return resp.Body, nil
}
