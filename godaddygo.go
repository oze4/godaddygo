package godaddygo

import (
	// "fmt"
	// "encoding/json"
	// "log"

	// "github.com/oze4/godaddygo/internal/endpoints/domains"
	"github.com/oze4/godaddygo/internal/core"
	// "github.com/oze4/godaddygo/internal/http"
)

// ClientInterface is what allows you to interact with the GoDaddy API
type ClientInterface interface {
	NewProduction() core.APIInterface
	NewDevelopment() core.APIInterface
}

// client implements ClientInterface
type client struct {
	*options
}

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func (c *client) NewProduction() core.APIInterface {
	return core.NewProductionAPI(c.apiKey, c.apiSecret)
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func (c *client) NewDevelopment() core.APIInterface {
	panic("The OTE (development) section of this library is under construction!")
	// return core.NewDevelopmentAPI()
}

// NewClient creates a new GoDaddy client.
func NewClient(o Options) ClientInterface {
	r := &client{&options{o.APIKey(), o.APISecret()}}
	return r
}

// Options hold the options for a new client
type Options interface {
	APIKey() string
	APISecret() string
}

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

// NewOptions returns a pointer to our client options
func NewOptions(apiKey, apiSecret string) Options {
	return &options{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

/*
// GetDomainInformation gets domain information
func (c Client) GetDomainInformation(d string) domains.DomainInformation {
	u := url.NewBuilder().Production().V1().Domain(d).Details()
	h := http.Client{
		APIKey:    c.Options.APIKey,
		APISecret: c.Options.APISecret,
	}

	r, e := h.Get(u)
	if e != nil {
		panic(e.Error())
	}

	var di domains.DomainInformation
	if err := json.Unmarshal(r, &di); err != nil {
		log.Fatal(err)
	}

	return di
}
*/
