package godaddygo

import (
	"github.com/oze4/godaddygo/pkg/session"
)

// DefaultClient defines behavior for a client
type DefaultClient interface {
	session.Interface
}

// NewDefaultClient creates a new default client
func NewDefaultClient(prod bool, key, secret string) DefaultClient {
	return &client{
		isProd:    prod,
		apiKey:    key,
		apiSecret: secret,
	}
}

// client implements Interface
type client struct {
	isProd    bool
	apiKey    string
	apiSecret string
}

// IsProduction determiens which base URL to use
func (c *client) IsProduction() bool {
	return c.isProd
}

// APIKey holds the api key
func (c *client) APIKey() string {
	return c.apiKey
}

// APISecret holds the API secret
func (c *client) APISecret() string {
	return c.apiSecret
}
