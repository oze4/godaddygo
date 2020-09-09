package gateway

import (
	"github.com/oze4/godaddygo/pkg/rest"
)

// NewAPIGateway ceates a new GoDaddy API gateway, which
// lets you select the API version
func NewAPIGateway(c *rest.Request) Interface {
	return &gateway{c}
}

// Interface allows you to select versions
type Interface interface {
	V1() V1Interface
	V2() V2Interface
}

// gateway implements Interface
type gateway struct {
	*rest.Request
}

func (g *gateway) V1() versions.V1Interface {
    g.APIVersion = "v1"
	return &v1{g.Request}
}

func (g *gateway) V2() error {
	panic("Not implemented")
}

// V1Interface targets version 1 of the GoDaddy API
type V1Interface interface {
	Domain(hostname string) domain.Interface
}

type v1 struct {
	*rest.Request
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v1) Domain(hostname string) domain.Interface {
	v.setTargetDomain(hostname)
	return &domain{v.meta}
}

// V2Interface targets version 1 of the GoDaddy API
type V2Interface interface {
    // v2 endpoints here
    // mirror v1
}

type v2 struct {
	*rest.Request
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v2) Domain(hostname string) domain.Interface {
	v.setTargetDomain(hostname)
	return &domain{v.meta}
}