package versions

import (
	"github.com/oze4/godaddygo/pkg/endpoints/domain"
)

// V1 targets version 1 of the GoDaddy API
type V1 interface {
	Domain(hostname string) domain.Interface
}

type v1 struct {
	*rest.Config
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v1) Domain(hostname string) domain.Interface {
	v.setTargetDomain(hostname)
	return &domain{v.meta}
}