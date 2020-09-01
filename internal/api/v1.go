package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// V1Getter makes embedding easier & more controllable
type V1Getter interface {
	V1() V1
}

// V1 holds all endpoints for version 1 of GoDaddy API
type V1 interface {
	DomainGetter
}

// v1 implements V1
type v1 struct {
	*http.Request
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v1) Domain(hostname string) Domain {
	v.URL = v.URL + "/v1"
	v.Host = hostname
	return &domain{v.Request}
}
