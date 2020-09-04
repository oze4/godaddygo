package godaddygo

import (
    "github.com/oze4/godaddygo/pkg/endpoints"
)

// Client holds required connection info
type Client interface {
	IsProduction() bool
	APIKey() string
    APISecret() string
	APIVersion() string
	// DNSDomainName() string
}

// client implements Client
type client struct {
	isProduction  bool
	apiKey        string
    apiSecret     string
    apiVersion    string          
	// dnsDomainName string
}

// NewClient creates a new Client
// If `isProd` is true, we use the "production" GoDaddy API (https://api.godaddy.com)
// If it is false we use the "development" (OTE) GoDaddy API (https://api-ote.godaddy.com)
func NewClient(isProd bool, key, secret string) Client {
    return &client{
        isProduction: isProd,
        apiKey: key,
        apiSecret: secret,
    }
}

// NewProductionAPI establishes a connection to the GoDaddy
// API endpoints
func NewProductionAPI(c Client) endpoints.Gateway {
    return &endpoints.Gateway{
        Session: &endpoints.Session{}
    }
}

func (g *client) IsProduction() bool {
    return g.isProduction
}

func (g *client) APIKey() string {
    return g.apiKey
}

func (g *client) APISecret() string {
    return g.apiSecret
}

func (g *client) APIVersion() string {
    return g.apiVersion
}