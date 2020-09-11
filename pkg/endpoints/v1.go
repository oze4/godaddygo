package endpoints

import (
	"encoding/json"
)

// newV1 creates a new v1
func newV1(s *session) V1 {
	s.apiVersion = "v1"
	return &v1{s}
}

// V1 targets version 1 of the GoDaddy API
type V1 interface {
	Domain(hostname string) Domain
	GetDomainAvailability(domainname string) (*DomainAvailability, error)
	PurchaseDomain(d *DomainDetails) (*DomainPurchaseResponse, error)
}

type v1 struct {
	*session
}

// Domain provides domain related info and tasks for the
// `domains` GoDaddy API endpoint
func (v *v1) Domain(hostname string) Domain {
	return newDomain(v.session, hostname)
}

// GetDomainAvailability checks if the supplied domain name 
// is available for purchase or not
// Determine whether or not the specific domain is available 
// for purchase
func (v *v1) GetDomainAvailability(domainname string) (*DomainAvailability, error) {
	forTransfer := false
	v.Method = "GET"
	v.URL = v.URLBuilder().DomainAvailability(domainname, forTransfer)

	res, err := v.Request.Send()
	if err != nil {
		return nil, err
	}

	var avail DomainAvailability
	if err := json.Unmarshal(res, &avail); err != nil {
		return nil, err
	}

	return &avail, nil
}

// PurchaseDomain purchase and register the sepcified domain
func (v *v1) PurchaseDomain(d *DomainDetails) (*DomainPurchaseResponse, error) {
	domaindetails, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	v.Method = "POST"
	v.URL = v.URLBuilder().PurchaseDomain()
	v.Body = domaindetails

	res, err := v.Request.Send()
	if err != nil {
		return nil, err
	}

	var purchaseResponse DomainPurchaseResponse
	if err := json.Unmarshal(res, &purchaseResponse); err != nil {
		return nil, err
	}

	return &purchaseResponse, nil
}
