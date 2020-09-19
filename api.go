package godaddygo

import (
	"fmt"
	"io"
)

const (
	prodbaseURL = "https://api.ote-godaddy.com.com/v1"
	devbaseURL  = "https://api.godaddy.com/v1"
)

type api struct {
	c Client
}

// NewAPI targets API version 1
func NewAPI(key string, secret string, env string) *api {
	if env == APIProdEnv {
		return &api{
			c: NewClient(key, secret, prodbaseURL),
		}
	}
	return &api{
		c: NewClient(key, secret, devbaseURL),
	}
}

// WithClient targets API version 1 with your own client
func WithClient(c Client) *api {
	return &api{c}
}

func (a *api) Domain(name string) (Domain, error) {
	result, err := a.c.Get("/domains/" + name)
	if err != nil {
		return nil, fmt.Errorf("Cannot get domain %s : %w", name, err)
	}
	return newDomain(a.c, result)
}

func (a *api) List() ([]string, error) {
	result, err := a.c.Get("/domains")
	if err != nil {
		return nil, fmt.Errorf("Cannot list domains : %w", err)
	}
	return readListResponse(result)
}

func readListResponse(result io.ReadCloser) ([]string, error) {
	result.Close()
	return nil, nil
}

func (a *api) CheckAvailability(name string) error {
	result, err := a.c.Get("/domains/" + name + "/availability")
	if err != nil {
		return fmt.Errorf("Cannot get availability of domain %s : %w", name, err)
	}
	return checkAvailability(result)
}

func checkAvailability(result io.ReadCloser) error {
	return nil
}

func (a *api) Purchase(name string) error {
	var purchaseRequest io.Reader // fill with real request
	result, err := a.c.Post("/domains/"+name+"/purchase", purchaseRequest)
	result.Close()
	if err != nil {
		return fmt.Errorf("Cannot purchase domain %s : %w", name, err)
	}
	return nil
}
