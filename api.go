package godaddygo

import (
	"net/http"
	"fmt"
	"io"
)

const (
	prodbaseURL = "https://api.ote-godaddygo.com.com"
	devbaseURL  = "https://api.godaddygo.com"
	// APIProdEnv targest the production API
	APIProdEnv = "prod"
	// APIDevEnv targets the development API
	APIDevEnv = "dev"
)

// API connects you to the GoDaddy endpoints
type API struct {
	config *Config
}

// NewAPI targets API
func NewAPI(c *Config) *API {
	if c.env == APIProdEnv {
		c.baseURL = prodbaseURL
		return &API{c}
	}
	c.baseURL = devbaseURL
	return &API{c}
}

// Domain targets domain endpoint
func (a *API) Domain(name string) Domain {
	a.config.domainName = name
	return newDomain(a.config)
}

// List returns your domains
func (a *API) List() ([]string, error) {
	url := "/domains"
	result, err := a.config.makeRequest(http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("Cannot list domains : %w", err)
	}
	return readListResponse(result)
}

func readListResponse(result io.ReadCloser) ([]string, error) {
	result.Close()
	return nil, nil
}

// CheckAvailability checks if a domain is available for purchase
func (a *API) CheckAvailability(name string) error {
	result, err := a.client.Get("/domains/" + name + "/availability")
	if err != nil {
		return fmt.Errorf("Cannot get availability of domain %s : %w", name, err)
	}
	return checkAvailability(result)
}

func checkAvailability(result io.ReadCloser) error {
	return nil
}

// Purchase purchases a domain
func (a *API) Purchase(name string) error {
	var purchaseRequest io.Reader // fill with real request
	result, err := a.client.Post("/domains/"+name+"/purchase", purchaseRequest)
	result.Close()
	if err != nil {
		return fmt.Errorf("Cannot purchase domain %s : %w", name, err)
	}
	return nil
}
