package version

import (
	"github.com/oze4/godaddygo/pkg/endpoint/domains"
)

// V1Interface targets version 1 of the GoDaddy API
type V1Interface interface {
	Domain(hostname string) domains.Domain
}

type v1Client struct {
	meta
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v v1Client) Domain(hostname string) domains.Domain {
	v.setTargetDomain(hostname)
	return &domain{v.meta}
}
