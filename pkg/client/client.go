package client

import (
	"github.com/oze4/godaddygo/internal/factory/url"
)

// Client is what allows you to interact with the GoDaddy API
type Client struct {
	Options    Options
	URLBuilder url.Builder
}

// NewClient creates a new GoDaddy client.
func NewClient(apiKey, apiSecret string) Client {
    opts := NewOptions(apiKey, apiSecret)
    return Client{Options: opts}
}

// Options hold the options for a new client
type Options struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

// NewOptions returns a pointer to our client options
func NewOptions(apiKey, apiSecret string) Options {
	return Options{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}