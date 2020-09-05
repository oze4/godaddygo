package godaddygo

import (
	"github.com/oze4/godaddygo/pkg/endpoints"
)

// NewClient creates a new Client
//  - If `prod` is true, we use the "production" GoDaddy API (https://api.godaddy.com)
//  - If it is false we use the "development" (OTE) GoDaddy API (https://api-ote.godaddy.com)
func NewClient(prod bool, key, secret string) endpoints.DefaultClient {
	return endpoints.NewDefaultClient(prod, key, secret)
}

// Connect connects you to GoDaddy API endpoints
func Connect(client endpoints.DefaultClient) endpoints.Gateway {
	return endpoints.NewConnection(client)
}