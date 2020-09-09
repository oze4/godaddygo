package endpoints

import (
	"encoding/json"

	"github.com/oze4/godaddygo/pkg/rest"
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
}

type v1 struct {
	*session
}

// Domain provides domain related info and tasks for the
// `domains` GoDaddy API endpoint
// This endpoint should be used for domains that you own
// (specifically, the `hostname` parameter)
func (v *v1) Domain(hostname string) Domain {
	return newDomain(v.session, hostname)
}

// GetDomainAvailability checks if the supplied domain name is available for purchase or not
func (v *v1) GetDomainAvailability(domainname string) (*DomainAvailability, error) {
    forTransfer := false
	req := &rest.Request{
		APIKey:    v.APIKey(),
		APISecret: v.APISecret(),
		URL:       v.URLBuilder().DomainAvailability(domainname, forTransfer),
		Method:    "GET",
	}

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	var avail DomainAvailability
	if err := json.Unmarshal(res, &avail); err != nil {
		return nil, err
	}

	return &avail, nil
}
