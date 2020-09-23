package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// newV1 is for internal convenience
func newV1(c *Config) *v1 {
	c.version = "v1"
	return &v1{c}
}

// v1 implements V1
type v1 struct {
	c *Config
}

// Domain targets domain endpoint
func (v *v1) Domain(name string) Domain {
	v.c.domainName = name
	return newDomain(v.c)
}

// ListDomains returns your domains
func (v *v1) ListDomains(ctx context.Context) ([]string, error) {
	url := "/domains"

	result, err := v.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, fmt.Errorf("Cannot list domains : %w", err)
	}

	return readListResponse(result)
}

// CheckAvailability checks if a domain is available for purchase
func (v *v1) CheckAvailability(ctx context.Context, name string) error {
	url := "/domains/" + name + "/availability"

	result, err := v.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return fmt.Errorf("Cannot get availability of domain %s : %w", name, err)
	}

	return checkAvailability(result)
}

// PurchaseDomain purchases a domain
func (v *v1) PurchaseDomain(ctx context.Context, dom DomainDetails) error {
	domBytes, err := json.Marshal(dom)
	if err != nil {
		return err
	}

	purchaseRequest := bytes.NewBuffer(domBytes)
	url := "/domains/" + v.c.domainName + "/purchase"

	if _, err := v.c.makeDo(ctx, http.MethodPost, url, purchaseRequest, 200); err != nil {
		return fmt.Errorf("Cannot purchase domain %s : %w", v.c.domainName, err)
	}

	return nil
}

// readListResponse reads http response when listing
func readListResponse(result io.ReadCloser) ([]string, error) {
	defer result.Close()
	return nil, nil
}

// checkAvailability reads the response for checking domain availability
func checkAvailability(result io.ReadCloser) error {
	defer result.Close()
	return nil
}
