package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
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

	response, err := makeDo(ctx, v.c, MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.listingDomains(err)
	}

	return readListDomainsResponse(response)
}

// CheckAvailability checks if a domain is available for purchase
func (v v1) CheckAvailability(ctx context.Context, name string, forTransfer bool) (DomainAvailability, error) {
	url := "/domains/available?domain=" + name + "&checkType=FAST&forTransfer=" + strconv.FormatBool(forTransfer)

	response, err := makeDo(ctx, v.c, MethodGet, url, nil, 200)
	if err != nil {
		return DomainAvailability{}, exception.checkingAvailability(err, name)
	}

	return readCheckAvailabilityResponse(response)
}

// PurchaseDomain purchases a domain
func (v v1) PurchaseDomain(ctx context.Context, dom DomainDetails) error {
	d, err := buildPurchaseDomainRequest(dom)
	if err != nil {
		return err
	}

	url := "/domains/" + v.c.domainName + "/purchase"
	if _, err := makeDo(ctx, v.c, MethodPost, url, d, 200); err != nil {
		return exception.purchasingDomain(err, v.c.domainName)
	}

	return nil
}

// readCheckAvailabilityResponse reads the response for checking domain availability
func readCheckAvailabilityResponse(r []byte) (DomainAvailability, error) {
	var availability DomainAvailability
	if err := json.Unmarshal(r, &availability); err != nil {
		return DomainAvailability{}, err
	}

	return availability, nil
}

// readListDomainsResponse reads http response when listing domains
func readListDomainsResponse(r []byte) ([]DomainSummary, error) {
	var domains []DomainSummary
	if err := json.Unmarshal(r, &domains); err != nil {
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