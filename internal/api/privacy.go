package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// PrivacyGetter makes embedding easier
type PrivacyGetter interface {
	Privacy() Privacy
}

// Privacy is a thing
type Privacy interface {
	Purchase() *http.Request
	Delete() *http.Request
}

// privacy implements Privacy
type privacy struct {
	*http.Request
}

func (p *privacy) Purchase() *http.Request {
	p.URL = p.URL + "/privacy/purchase"
	p.Method = "POST"
	return p.Request
}

func (p *privacy) Delete() *http.Request {
	p.URL = p.URL + "/privacy"
	p.Method = "DELETE"
	return p.Request
}
