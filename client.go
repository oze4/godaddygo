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

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func (c *client) NewProduction() api.Gateway {
	return api.NewProductionAPI(c.apiKey, c.apiSecret)
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func (c *client) NewDevelopment() api.Gateway {
	panic("The OTE (development) section of this library is under construction!")
	// return core.NewDevelopmentAPI()
}

// NewClient creates a new GoDaddy client.
func NewClient(o *options) Client {
	r := &client{options: o}
	return r
}
