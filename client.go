package godaddygo

import (
	"github.com/oze4/godaddygo/internal/api"
)

// Client is what allows you to interact with the GoDaddy API
type Client interface {
	NewProduction() api.Gateway
	NewDevelopment() api.Gateway
}

// client implements Client
type client struct {
	*options
}

// NewClient creates a new GoDaddy client.
func NewClient(o *options) Client {
	return &client{options: o}
}

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func (c *client) NewProduction() api.Gateway {
	api.
	/*
	return api.NewGatewayWithRequest(&request{
		apiKey:    key,
		apiSecret: secret,
		url:       "https://api.godaddy.com",
	})
	*/
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func (c *client) NewDevelopment() api.Gateway {
	panic("The OTE (development) section of this library is under construction!")
	// return api.NewDevelopment(c.apiKey, c.apiSecret)
}

// NewRequest returns a new request object
func (c *client) NewRequest() api.Request {
    return &request{}
}
