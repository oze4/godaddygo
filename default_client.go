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
		Prod: prod,
		apiKey: key,
		apiSecret: secret,
	}
}

// client implements Interface
type client struct {
    Prod bool
	apiKey       string
    apiSecret    string
}

// IsProduction determiens which base URL to use
func (c *client) IsProduction() bool {
	return c.Prod
}

// APIKey holds the api key
func (c *client) APIKey() string {
	return c.apiKey
}

// APISecret holds the API secret
func (c *client) APISecret() string {
	return c.apiSecret
}

/*
// SetAPIKey sets the API key
func (c *client) SetAPIKey(k string) {
    c.apiKey = k
}

// SetAPISecret sets the API secret
func (c *client) SetAPISecret(s string) {
    c.apiSecret = s
}

// SetIsProduction helps determine which base URL to use
func (c *client) SetIsProduction(p bool) {
    c.Prod = p
}
*/