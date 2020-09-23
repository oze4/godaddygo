package godaddygo

import (
    "net/http"
)

// NewConfig creates a default config, allowing you to choose Production or
// Development GoDaddy API via `APIProdEnv` or `APIDevEnv` constants
func NewConfig(key, secret, env string) *Config {
	return newConfig(key, secret, env)
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
	c := newConfig(key, secret, APIProdEnv)
    g, _ := New(c)
    return g
}

// NewDevelopment targets GoDaddy development API
func NewDevelopment(key string, secret string) Gateway {
    c := newConfig(key, secret, APIDevEnv)
    g, _ := New(c)
	return g
}

// WithClient returns a Gateway using your own `*http.Client`
func WithClient(conf *Config, client *http.Client) Gateway {
    conf.client = client
    g, _ := New(conf)
    return g
}
