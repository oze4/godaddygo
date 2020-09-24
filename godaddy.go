package godaddygo

import (
	"net/http"
)

// NewConfig creates a new config
func NewConfig(key, secret, env string) *Config {
	return newConfig(key, secret, env)
}

// NewProduction targets GoDaddy production API
func NewProduction(key string, secret string) (Gateway, error) {
	c := newConfig(key, secret, APIProdEnv)
	return new(c)
}

// NewDevelopment targets GoDaddy development API
func NewDevelopment(key string, secret string) (Gateway, error) {
	c := newConfig(key, secret, APIDevEnv)
	return new(c)
}

// WithClient returns a Gateway using your own `*http.Client`
func WithClient(client *http.Client, config *Config) (Gateway, error) {
	config.client = client
	return new(config)
}

// new returns a new Gateway based upon a config
// Also sets the `baseURL` based upon `env`
func new(c *Config) (Gateway, error) {
	switch c.env {
	case APIProdEnv:
		c.baseURL = prodbaseURL
	case APIDevEnv:
		c.baseURL = devbaseURL
	default:
		return nil, exception.invalidAPIEnv(nil)
	}

	return newGateway(c), nil
}
