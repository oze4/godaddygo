package godaddygo

import (
	"errors"
)

// PrivacyGetter makes embedding easier
type PrivacyGetter interface {
	Privacy() Privacy
}

// Privacy is a thing
type Privacy interface {
	Purchase() error
	Delete() error
}

// privacy implements Privacy
type privacy struct {
	connectionBridge
}

func (p *privacy) Purchase() error {
	// p.URL = p.URL + "/privacy/purchase"
	// p.Method = "POST"
	// return p.Request
	return errors.New("Not implemented")
}

func (p *privacy) Delete() error {
	// p.URL = p.URL + "/privacy"
	// p.Method = "DELETE"
	// return p.Request
	return errors.New("Not implemented")
}
