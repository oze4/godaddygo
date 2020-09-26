package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// newV1 is for internal convenience
func newV1(c *Config) v1 {
	c.version = APIVersion1
	return v1{c}
}

// v1 implements V1
type v1 struct {
	c *Config
}

// Domain targets domain endpoint
func (v v1) Domain(name string) Domain {
	v.c.domainName = name
	return newDomain(v.c)
}

// ListDomains returns your domains
func (v v1) ListDomains(ctx context.Context) ([]DomainSummary, error) {
	url := "/domains"
	result, err := v.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.listingDomains(err)
	}

	return readListDomainsResponse(result)
}

// CheckAvailability checks if a domain is available for purchase
func (v v1) CheckAvailability(ctx context.Context, name string, forTransfer bool) (DomainAvailability, error) {
	url := "/domains/available?domain=" + name + "&checkType=FAST&forTransfer=" + strconv.FormatBool(forTransfer)
	result, err := v.c.makeDo(ctx, http.MethodGet, url, nil, 200)
	if err != nil {
		return DomainAvailability{}, exception.checkingAvailability(err, name)
	}

	return readCheckAvailabilityResponse(result)
}

// PurchaseDomain purchases a domain
func (v v1) PurchaseDomain(ctx context.Context, dom DomainDetails) error {
	url := "/domains/" + v.c.domainName + "/purchase"
	d, err := buildPurchaseDomainRequest(dom)
	if err != nil {
		return err
	}

	if _, err := v.c.makeDo(ctx, http.MethodPost, url, d, 200); err != nil {
		return exception.purchasingDomain(err, dom.Domain)
	}
	return nil
}

// readCheckAvailabilityResponse reads the response for checking domain availability
func readCheckAvailabilityResponse(result io.ReadCloser) (DomainAvailability, error) {
	defer result.Close()
	content, err := bodyToBytes(result)
	if err != nil {
		return DomainAvailability{}, err
	}

	var availability DomainAvailability
	if err := json.Unmarshal(content, &availability); err != nil {
		return DomainAvailability{}, err
	}

	return availability, nil
}

// readListDomainsResponse reads http response when listing domains
func readListDomainsResponse(result io.ReadCloser) ([]DomainSummary, error) {
	defer result.Close()
	content, err := bodyToBytes(result)
	if err != nil {
		return []DomainSummary{}, err
	}

	var domains []DomainSummary
	if err := json.Unmarshal(content, &domains); err != nil {
		return []DomainSummary{}, err
	}

	return domains, nil
}

// buildPurchaseDomainRequest marshals domain details object and returns it as a byte.Buffer
func buildPurchaseDomainRequest(dom DomainDetails) (*bytes.Buffer, error) {
	domBytes, err := json.Marshal(dom)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(domBytes), nil
}
