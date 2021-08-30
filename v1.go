package godaddygo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/oze4/godaddygo/internal/exception"
)

// V1 knows how to interact with GoDaddy Gateway version 1
type V1 interface {
	Domain(name string) Domain
	ListDomains(ctx context.Context) ([]DomainSummary, error)
	CheckAvailability(ctx context.Context, name string, forTransfer bool) (DomainAvailability, error)
	PurchaseDomain(ctx context.Context, p PurchaseRequest) (PurchaseReceipt, error)
}

// newV1 is for internal convenience
func newV1(config *Config) v1 {
	config.version = APIVersion1
	return v1{config}
}

// v1 implements V1
type v1 struct {
	config *Config
}

func (v v1) Consent(agreedAt, agreedBy string, privacy, forTransfer bool, tlds []string) Consent {
	return newConsent(agreedAt, agreedBy, privacy, forTransfer, tlds)
}

// Domain targets domain endpoint
func (v v1) Domain(name string) Domain {
	v.config.domainName = name
	return newDomain(v.config)
}

// ListDomains returns your domains
func (v v1) ListDomains(ctx context.Context) ([]DomainSummary, error) {
	url := "/domains"
	response, err := makeDo(ctx, v.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return nil, exception.ListingDomains(err)
	}
	return readListDomainsResponse(response)
}

// CheckAvailability checks if a domain is available for purchase
func (v v1) CheckAvailability(ctx context.Context, name string, forTransfer bool) (DomainAvailability, error) {
	url := "/domains/available?domain=" + name + "&checkType=FAST&forTransfer=" + strconv.FormatBool(forTransfer)
	response, err := makeDo(ctx, v.config, http.MethodGet, url, nil, 200)
	if err != nil {
		return DomainAvailability{}, exception.CheckingAvailability(err, name)
	}
	return readCheckAvailabilityResponse(response)
}

// PurchaseDomain purchases a domain
func (v v1) PurchaseDomain(ctx context.Context, p PurchaseRequest) (PurchaseReceipt, error) {
	d, err := buildPurchaseDomainRequest(p)
	if err != nil {
		return PurchaseReceipt{}, err
	}
	url := "/domains/purchase"
	if _, err := makeDo(ctx, v.config, http.MethodPost, url, d, 200); err != nil {
		return PurchaseReceipt{}, exception.PurchasingDomain(err, p.Domain)
	}
	return PurchaseReceipt{}, nil
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
func buildPurchaseDomainRequest(p PurchaseRequest) (*bytes.Buffer, error) {
	if !p.IsValidPeriod() {
		return nil, fmt.Errorf("invalid Period provided : min 1 max 10")
	}
	domBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(domBytes), nil
}
