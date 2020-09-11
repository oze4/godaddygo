package endpoints

import (
    "encoding/json"
    "time"
)

func newDomains(s *session) Domains {
	return &domains{s}
}

// DomainsGetter simplifies embedding
type DomainsGetter interface {
	Domains() Domains
}

// Domains knows how to interact with domains you may or may not
// own but want to perform tasks on.
//  - Check if a domain is available for purchase
//  - Purchase a domin
//  - List all of the domains you own
// etc...
type Domains interface {
	CheckAvailability(domainname string) (*DomainAvailability, error)
	My() (*[]GoDaddyDomain, error)
	Purchase(domaindetails *DomainDetails) (*DomainPurchaseResponse, error)
}

// domains (plural) implements Domains (plural)
type domains struct {
	*session
}

// My lists all domains *you own* or have the ability to manage
func (d *domains) My() (*[]GoDaddyDomain, error) {
	d.Method = "GET"
	d.URL = d.URLBuilder().GetMyDomains()

	res, err := d.Request.Send()
	if err != nil {
		return nil, err
	}

	var mydomains []GoDaddyDomain
	if err := json.Unmarshal(res, &mydomains); err != nil {
		return nil, err
	}

	return &mydomains, nil
}

// CheckAvailability determine whether or not the specific domain is available
// for purchase
func (d *domains) CheckAvailability(domainname string) (*DomainAvailability, error) {
	forTransfer := false
	d.Method = "GET"
	d.URL = d.URLBuilder().DomainAvailability(domainname, forTransfer)

	res, err := d.Request.Send()
	if err != nil {
		return nil, err
	}

	var avail DomainAvailability
	if err := json.Unmarshal(res, &avail); err != nil {
		return nil, err
	}

	return &avail, nil
}

// Purchase purchase and register the sepcified domain
func (d *domains) Purchase(domaindetails *DomainDetails) (*DomainPurchaseResponse, error) {
	domdetails, err := json.Marshal(domaindetails)
	if err != nil {
		return nil, err
	}

	d.Method = "POST"
	d.URL = d.URLBuilder().PurchaseDomain()
	d.Body = domdetails

	res, err := d.Request.Send()
	if err != nil {
		return nil, err
	}

	var purchaseResponse DomainPurchaseResponse
	if err := json.Unmarshal(res, &purchaseResponse); err != nil {
		return nil, err
	}

	return &purchaseResponse, nil
}

// GoDaddyDomain is what gets returned when listing all of
// your domains
type GoDaddyDomain struct {
	CreatedAt              time.Time
	Domain                 string
	DomainID               int
	ExpirationProtected    bool
	Expires                time.Time
	ExposeWhois            bool
	HoldRegistrar          bool
	Locked                 bool
	NameServers            interface{} // []string
	Privacy                bool
	RenewAuto              bool
	Renewable              bool
	Status                 string
	TransferAwayEligibleAt time.Time
	TransferProtected      bool
}
