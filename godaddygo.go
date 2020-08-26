package godaddygo

import (
    godaddy "github.com/oze4/godaddygo/pkg/client"
)

// NewClient creates a new GoDaddy client.
func NewClient(apiKey, apiSecret string) godaddy.Client {
    return godaddy.NewClient(apiKey, apiSecret)
}
