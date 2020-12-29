package godaddygo

import (
	"net/http"
	"time"
)

const (
	// MaxHTTPClientTimeout in seconds is the defalt max http client timeout
	// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	// 1 minute by default
	MaxHTTPClientTimeout = time.Second * 60
)

// NewConfig creates a new config
func NewConfig(key, secret string, env APIEnv) *Config {
	return &Config{
		client: &http.Client{
			Timeout: MaxHTTPClientTimeout,
		},
		key:    key,
		secret: secret,
		env:    env,
	}
}

// NewProduction targets GoDaddy production API
func NewProduction(key string, secret string) (API, error) {
	c := NewConfig(key, secret, APIProdEnv)
	return newAPI(c)
}

// NewDevelopment targets GoDaddy development API
func NewDevelopment(key string, secret string) (API, error) {
	c := NewConfig(key, secret, APIDevEnv)
	return newAPI(c)
}

// WithClient returns API using your own `*http.Client`
func WithClient(client *http.Client, config *Config) (API, error) {
	config.client = client
	return newAPI(config)
}
