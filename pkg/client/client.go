package client

import (
	"encoding/json"
	"log"

	"github.com/oze4/godaddygo/internal/endpoints/domains"
	"github.com/oze4/godaddygo/internal/factory/url"
	"github.com/oze4/godaddygo/internal/http"
)

// Client is what allows you to interact with the GoDaddy API
type Client struct {
	Options    Options
	URLBuilder url.Builder // I'm not sure if we should expose this or not..
}

// NewClient creates a new GoDaddy client.
func NewClient(apiKey, apiSecret string) Client {
	opts := NewOptions(apiKey, apiSecret)
	return Client{Options: opts}
}

// Options hold the options for a new client
type Options struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

// NewOptions returns a pointer to our client options
func NewOptions(apiKey, apiSecret string) Options {
	return Options{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

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
