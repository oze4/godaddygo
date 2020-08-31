package godaddygo

import (
	"github.com/oze4/godaddygo/internal/api"
	"github.com/oze4/godaddygo/pkg/http"
)

// NewProductionAPI targets GoDaddy's production API (https://api.godaddy.com)
func NewProductionAPI(o Options) api.Gateway {
	return api.InitProduction(&http.Request{
		APIKey:    o.APIKey(),
		APISecret: o.APISecret(),
	})
}

// NewDevelopmentAPI targets GoDaddy's development API (https://api.ote-godaddy.com)
func NewDevelopmentAPI(o Options) api.Gateway {
	return api.InitDevelopment(&http.Request{
		APIKey:    o.APIKey(),
		APISecret: o.APISecret(),
	})
}

// Options hold the options for a new client
type Options interface {
	APIKey() string
	APISecret() string
}

// options implements Options
type options struct {
	apiKey    string
	apiSecret string
}

func (o *options) APIKey() string {
	return o.apiKey
}

func (o *options) APISecret() string {
	return o.apiSecret
}

// NewOptions returns new options
func NewOptions(apikey, apisecret string) Options {
	return &options{apikey, apisecret}
}
