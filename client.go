package godaddygo

import (
	"github.com/oze4/godaddygo/pkg/endpoints"
	"github.com/oze4/godaddygo/pkg/session"
)

// NewClient creates a new Client
// If `isProd` is true, we use the "production" GoDaddy API (https://api.godaddy.com)
// If it is false we use the "development" (OTE) GoDaddy API (https://api-ote.godaddy.com)
func NewClient(isProd bool, key, secret string) Client {
	return &client{
		isProduction: isProd,
		apiKey:       key,
		apiSecret:    secret,
	}
}

// Connect connects you to the endpoints
func Connect(c Client) endpoints.Gateway {
	return endpoints.NewConnection(c)
}

// Client defines behavior for a client
type Client interface {
	session.Interface
}

// client implements Client
type client struct {
	isProduction bool
	apiKey       string
	apiSecret    string
	apiVersion   string
}

func (c *client) IsProduction() bool {
	return c.isProduction
}

func (c *client) APIKey() string {
	return c.apiKey
}

func (c *client) APISecret() string {
	return c.apiSecret
}

func (c *client) APIVersion() string {
	return c.apiVersion
}
