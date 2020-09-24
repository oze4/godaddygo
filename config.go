package godaddygo

import (
	"context"
	"io"
	"net/http"
)

func newConfig(key, secret, env string) *Config {
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

// makeDo makes an http.Request and sends it
func (c *Config) makeDo(ctx context.Context, method string, path string, body io.Reader, expectedStatus int) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.baseURL+"/"+c.version+path, body)
	if err != nil {
		return nil, exception.creatingNewRequest(err)
	}

	req.WithContext(ctx)

	req.Header.Set("Authorization", "sso-key "+c.key+":"+c.secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, exception.sendingRequest(err)
	}

	if resp.StatusCode != expectedStatus {
		return resp.Body, exception.invalidStatusCode(expectedStatus, resp.StatusCode, err)
	}

	return resp.Body, nil
}
