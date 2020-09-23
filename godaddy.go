package godaddygo

import (
    "net/http"
)

// NewConfig creates a default config, allowing you to choose Production or
// Development GoDaddy API via `APIProdEnv` or `APIDevEnv` constants
func NewConfig(key, secret, env string) *Config {
	return newDefaultConfig(key, secret, env)
}

// WithClient returns a Gateway using your own `*http.Client`
func WithClient(conf *Config, client *http.Client) Gateway {
    conf.withClient(client)
    return newGateway(conf)
}

// New returns a new Gateway based upon a config
func New(c *Config) (Gateway, error) {
	switch c.env {
	case APIProdEnv:
		c.baseURL = prodbaseURL
	case APIDevEnv:
		c.baseURL = devbaseURL
	default:
		return nil, ErrorWrongAPIEnv
    }
    
	return newGateway(c), nil
}

// NewProduction targets GoDaddy production API
func NewProduction(key string, secret string) Gateway {
	c := newDefaultConfig(key, secret, APIProdEnv)
	return newGateway(c)
}

// NewDevelopment targets GoDaddy development API
func NewDevelopment(key string, secret string) Gateway {
	c := newDefaultConfig(key, secret, APIDevEnv)
	return newGateway(c)
}
