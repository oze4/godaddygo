package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// API connects you to the GoDaddy endpoints
type API struct {
	c *Config
}

// NewAPI uses a config to connect you to the GoDaddy API
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
	a.c.domainName = name
	return newDomain(a.c)
}

// List returns your domains
func (a *API) List(ctx context.Context) ([]string, error) {
	url := "/domains"

	result, err := a.c.makeDo(ctx, http.MethodGet, url, nil, 200)
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
func (a *API) CheckAvailability(ctx context.Context, name string) error {
	url := "/domains/" + name + "/availability"

	result, err := a.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return fmt.Errorf("Cannot get availability of domain %s : %w", name, err)
	}

	return checkAvailability(result)
}

func checkAvailability(result io.ReadCloser) error {
	return nil
}

// Purchase purchases a domain
func (a *API) Purchase(ctx context.Context, dom DomainDetails) error {
	domBytes, err := json.Marshal(dom)
	if err != nil {
		return err
	}

	purchaseRequest := bytes.NewBuffer(domBytes)
	url := "/domains/" + a.c.domainName + "/purchase"

	if _, err := a.c.makeDo(ctx, http.MethodPost, url, purchaseRequest, 200); err != nil {
		return fmt.Errorf("Cannot purchase domain %s : %w", a.c.domainName, err)
	}

	return nil
}
