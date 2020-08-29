package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// Gateway allows you to select different versions of the GoDaddy API
type Gateway interface {
	V1Getter
}

// versions implements APIInterface
type gateway struct {
	*http.Request
}

// V1 returns the V1 section of the GoDaddy API
func (a *gateway) V1() V1 {
	a.URL = a.URL + "/v1"
	return &v1{a.Request}
}

