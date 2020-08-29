package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// V1Getter does a thing
type V1Getter interface {
	V1() V1
}

// V1 does a thing
type V1 interface {
	DomainGetter
}

type v1 struct {
	*http.Request
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v1) Domain(name string) Domain {
	v.Host = name
	v.URL = v.URL + "/domains/" + name
	return &domain{v.Request}
}
