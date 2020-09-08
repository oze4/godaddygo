package rest

import (
    "github.com/oze4/godaddygo/pkg/gateway"
)

// Client holds request config
type Client struct {
    config Config
}

// NewRESTClient creates a new REST client
func NewRESTClient(apikey, apisecret string, isproduction bool) *Client {
    return &Client{
        config: Config{
            APIKey: apikey,
            APISecret: apisecret,
            IsProduction: isproduction,
        },
    }
}

func t() {
    gateway.
}