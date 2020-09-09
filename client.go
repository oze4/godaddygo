package godaddygo

import (
	"github.com/oze4/godaddygo/pkg/client"
	"github.com/oze4/godaddygo/pkg/endpoints"
)

// NewClient creates a new Client
//  - If `prod` is true, we use the "production" GoDaddy API (https://api.godaddy.com)
//  - If it is false we use the "development" (OTE) GoDaddy API (https://api-ote.godaddy.com)
func NewClient(prod bool, key, secret string) client.Interface {
	return client.NewClient(key, secret, prod)
}

// Connect connects you to GoDaddy API endpoints
func Connect(godaddygoClient client.Interface) endpoints. {
	return endpoints.NewConnection(godaddygoClient)
}
