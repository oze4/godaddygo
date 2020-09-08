package domain

import (
	"errors"

	"github.com/oze4/godaddygo/pkg/rest"
)

// PrivacyGetter makes embedding easier
type PrivacyGetter interface {
	Privacy() PrivacyInterface
}

// PrivacyInterface lets you remove or purchase domain
// privacy protection
type PrivacyInterface interface {
	Purchase() error
	Delete() error
}

// privacy implements Privacy
type privacy struct {
	*rest.Config
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
