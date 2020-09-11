package endpoints

import (
	"encoding/json"
)

func newPrivacy(s *session) Privacy {
	return &privacy{s}
}

// PrivacyGetter makes embedding easier
type PrivacyGetter interface {
	Privacy() Privacy
}

// Privacy lets you remove or purchase domain
// privacy protection
type Privacy interface {
	Purchase(c *Consent) error
	Remove() error
}

// privacy implements Privacy
type privacy struct {
	*session
}

func (p *privacy) Purchase(c *Consent) error {
	purchaseconsent, err := json.Marshal(c)
	if err != nil {
		return err
	}

	p.Method = "POST"
	p.URL = p.URLBuilder().Domain(p.domainName).PurchasePrivacy()
	p.Body = purchaseconsent

	if _, err := p.Request.Send(); err != nil {
		return err
	}

	return nil
}

func (p *privacy) Remove() error {
	p.Method = "DELETE"
	p.URL = p.URLBuilder().Domain(p.domainName).RemovePrivacy()

	if _, err := p.Request.Send(); err != nil {
		return err
	}

	return nil
}
