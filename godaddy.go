package godaddygo

import (
	"net/http"
)

// NewConfig creates a new config
func NewConfig(key, secret, env string) *Config {
	return &Config{client: http.DefaultClient, key: key, secret: secret, env: env}
}

// NewProduction targets GoDaddy production API
func NewProduction(key string, secret string) (API, error) {
	return newAPI(NewConfig(key, secret, APIProdEnv))
}

// NewDevelopment targets GoDaddy development API
func NewDevelopment(key string, secret string) (API, error) {
	return newAPI(NewConfig(key, secret, APIDevEnv))
}

// WithClient returns API using your own `*http.Client`
func WithClient(client *http.Client, config *Config) (API, error) {
	config.client = client
	return newAPI(config)
}
