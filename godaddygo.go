package godaddygo

import (
    godaddy "github.com/oze4/godaddygo/pkg/client"
)

// NewClient creates a new GoDaddy client.
func NewClient(apiKey, apiSecret, apiMode string) godaddy.Client {
    opts := godaddy.NewOptions(apiKey, apiSecret, apiMode)
    return godaddy.Client{Options: opts}
}
