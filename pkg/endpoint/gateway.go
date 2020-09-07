package endpoint

import (
    "github.com/oze4/godaddygo/pkg/rest"
    "github.com/oze4/godaddygo/pkg/endpoint/versions"
)

// Gateway allows you to select versions
type Gateway interface {
	V1() V1
}

// gateway implements Gateway
type gateway struct {
	rest.Config
}

func (g *gateway) V1() versions.V1Interface {
    g.APIVersion = "v1"
	return &v1{g.}
}