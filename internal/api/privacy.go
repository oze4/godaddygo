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
}

// privacy implements Privacy
type privacy struct {
	*http.Request
}

func (p *privacy) Purchase() *http.Request {
	p.Method = "POST"
	p.URL = p.URL + "privacy/purchase"
	return p.Request
}

func (p *privacy) Delete() *http.Request {
	p.Method = "DELETE"
	p.URL = p.URL + "/privacy"
	return p.Request
}