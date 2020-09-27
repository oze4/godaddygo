package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

	response, err := v.c.makeDo(ctx, MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.listingDomains(err)
	}

	return readListDomainsResponse(response)
}

// CheckAvailability checks if a domain is available for purchase
func (v v1) CheckAvailability(ctx context.Context, name string, forTransfer bool) (DomainAvailability, error) {
	url := "/domains/available?domain=" + name + "&checkType=FAST&forTransfer=" + strconv.FormatBool(forTransfer)

	response, err := v.c.makeDo(ctx, MethodGet, url, nil, 200)
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

	return readPurchaseDomainResponse(ctx, v.c, d)
}

// readCheckAvailabilityResponse reads the response for checking domain availability
func readCheckAvailabilityResponse(result io.ReadCloser) (DomainAvailability, error) {
	var availability DomainAvailability
	if err := readResponseTo(result, &availability); err != nil {
		return DomainAvailability{}, err
	}

	return availability, nil
}

// readListDomainsResponse reads http response when listing domains
func readListDomainsResponse(result io.ReadCloser) ([]DomainSummary, error) {
	var domains []DomainSummary
	if err := readResponseTo(result, &domains); err != nil {
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

func readPurchaseDomainResponse(ctx context.Context, conf *Config, body *bytes.Buffer) error {
	url := "/domains/" + conf.domainName + "/purchase"
	if _, err := conf.makeDo(ctx, MethodPost, url, body, 200); err != nil {
		return exception.purchasingDomain(err, conf.domainName)
	}

	return nil
}
