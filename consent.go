package godaddygo

import "fmt"

// Gives consent to purchase a domain

/**
MIT License

Copyright (c) 2021-CURRENT Matt Oestreich

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.



PURCHASE A DOMAIN AT YOUR OWN RISK!!!
*/

// Consent is needed when purhasing a domain
type Consent interface {
	Agree(yesIAmSure bool) error
	AgreedAt() string
	AgreedBy() string
	AgreementKeys() ([]string, error)
}

// NewConsent begins the agreement process
// Per GoDaddy, `agreedAt` must be in `iso-datetime` format (RFC3339)
func newConsent(agreedAt, agreedBy string, privacy, forTransfer bool, tlds []string) Consent {
	return &consent{
		isAgreed: false,
		agreedAt: agreedAt,
		agreedBy: agreedBy,
		privacy: privacy,
		forTransfer: forTransfer,
		tlds: tlds,
	}
}

type consent struct {
	agreedAt      string
	agreedBy      string
	agreementKeys []string
	isAgreed      bool
	privacy       bool
	forTransfer   bool
	tlds []string
}

func (c *consent) Agree(yesIAmSure bool) error {
	if !yesIAmSure {
		return fmt.Errorf("without consent we cannot purchase a domain : you answered '%t' when asked to confirm consent", yesIAmSure)
	}
	//TODO: get (and then set) agreement keys, finally return them
	// https://developer.godaddy.com/doc/endpoint/domains#/v1/purchase
	c.isAgreed = true
	return nil
}

func (c *consent) AgreedAt() string {
	return c.agreedAt
}

func (c *consent) AgreedBy() string {
	return c.agreedBy
}

func (c *consent) AgreementKeys() ([]string, error) {
	if !c.isAgreed {
		return nil, fmt.Errorf("you must call `Consent.Agree(bool)` before we can get your agreement documents")
	}
	return c.agreementKeys, nil
}

func buildDomainAgreementKeysRequest(c consent, cfg *Config) {

}
