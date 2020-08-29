package godaddygo

import (
	"github.com/oze4/godaddygo/internal/api"
	"github.com/oze4/godaddygo/pkg/http"
)

// Client is what allows you to interact with the GoDaddy API
type Client interface {
	NewProduction() api.Gateway
	NewDevelopment() api.Gateway
}

// client implements Client
type client struct {
	*Options
}

// NewClient creates a new GoDaddy client.
func NewClient(opts Options) Client {
	return &client{&opts}
}

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func (c *client) NewProduction() api.Gateway {
	h := &http.Request{
		APIKey:    c.APIKey,
		APISecret: c.APISecret,
		URL:       "https://api.godaddy.com",
	}

	return api.InitProduction(h)
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func (c *client) NewDevelopment() api.Gateway {
	panic("The OTE (development) section of this library is under construction!")
	// return api.InitDevelopment(c.apiKey, c.apiSecret)
}

// NewRequest returns a new request object
func (c *client) NewRequest() *http.Request {
	return &http.Request{}
}
