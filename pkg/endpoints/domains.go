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

// Domains knows how to interact with domains you may not own
//
//  - check if a domain is available for purchase
//  - purchase a domain
//
// Domains is also useful for when you don't want to target
// a specific domain
//
//  - list all domains you own
//  - etc..
type Domains interface {
	ContactsGetter
	PrivacyGetter
	RecordsGetter
	CheckAvailability(domainname string) (*DomainAvailability, error)
	GetDetails(domainname string) (*DomainDetails, error)
	List() (*[]DomainDetails, error)
	Purchase(domaindetails *DomainDetails) (*DomainPurchaseResponse, error)
}

type domains struct {
	*session
}

func (d *domains) Records(domainname string) Records {
	return newRecords(d.session, domainname)
}

// Contacts builds out the contacts piece of the URL
func (d *domains) Contacts(domainname string) Contacts {
	return newContacts(d.session, domainname)
}

// Privacy builds out the privacy piece of the URL
func (d *domains) Privacy(domainname string) Privacy {
	return newPrivacy(d.session, domainname)
}

func (d *domains) List() (*[]DomainDetails, error) {
	d.Method = "GET"
	d.URL = d.URLBuilder().GetMyDomains()

	res, err := d.Request.Send()
	if err != nil {
		return nil, err
	}

	var mydomains []DomainDetails
	if err := json.Unmarshal(res, &mydomains); err != nil {
		return nil, err
	}

	return &mydomains, nil
}

// GetDetails gets info on a domain
func (d *domains) GetDetails(domainname string) (*DomainDetails, error) {
	d.Method = "GET"
	d.URL = d.URLBuilder().Domain(domainname).String()

	res, err := d.Request.Send()
	if err != nil {
		return nil, err
	}

	var details DomainDetails
	if err := json.Unmarshal(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
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

// DomainDetails holds information about a GoDaddy domain.
// This is the response when you `GET` info about a domain.
type DomainDetails struct {
	AuthCode               string        `json:"authCode,omitempty"`
	ContactAdmin           Contact       `json:"contactAdmin,omitempty"`
	ContactBilling         Contact       `json:"contactBilling,omitempty"`
	ContactRegistrant      Contact       `json:"contactRegistrant,omitempty"`
	ContactTech            Contact       `json:"contactTech,omitempty"`
	CreatedAt              time.Time     `json:"createdAt,omitempty"`
	DeletedAt              time.Time     `json:"deletedAt,omitempty"`
	TransferAwayEligibleAt time.Time     `json:"transferAwayEligibleAt,omitempty"`
	Domain                 string        `json:"domain,omitempty"`
	DomainID               int           `json:"domainId,omitempty"`
	ExpirationProtected    bool          `json:"expirationProtected,omitempty"`
	Expires                time.Time     `json:"expires,omitempty"`
	ExposeWhois            bool          `json:"exposeWhois,omitempty"`
	HoldRegistrar          bool          `json:"holdRegistrar,omitempty"`
	Locked                 bool          `json:"locked,omitempty"`
	NameServers            []string      `json:"nameServers,omitempty"`
	Privacy                bool          `json:"privacy,omitempty"`
	RenewAuto              bool          `json:"renewAuto,omitempty"`
	RenewDeadline          time.Time     `json:"renewDeadline,omitempty"`
	Status                 string        `json:"status,omitempty"`
	SubAccountID           string        `json:"subaccountId,omitempty"`
	TransferProtected      bool          `json:"transferProtected,omitempty"`
	Verifications          Verifications `json:"verifications,omitempty"`
}

// Verifications holds verification info about a domain.
type Verifications struct {
	DomainName DomainName `json:"domainName,omitempty"`
	RealName   RealName   `json:"realName,omitempty"`
}

// RealName holds verifications real name info
type RealName struct {
	Status string `json:"status,omitempty"`
}

// DomainName holds verification domain name info
type DomainName struct {
	Status string `json:"status,omitempty"`
}

// DomainAvailability holds data about domain availability
type DomainAvailability struct {
	Available  bool   `json:"available,omitempty"`
	Currency   string `json:"currency,omitempty"`
	Definitive bool   `json:"definitive,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Period     int    `json:"period,omitempty"`
	Price      int    `json:"price,omitempty"`
}

// DomainPurchaseResponse is GoDaddy's response to purchasing a domain
type DomainPurchaseResponse struct {
	Currency  string `json:"currency,omitempty"`
	ItemCount int    `json:"itemCount,omitempty"`
	OrderID   int    `json:"orderId,omitempty"`
	Total     int    `json:"total,omitempty"`
}

// Consent is required when purhasing domain privacy
type Consent struct {
	AgreedAt      time.Time
	AgreedBy      string
	AgreementKeys []string // No idea what this is, need to dig into that
}

// Update updates a domain
// func (d *domain) Update(body []byte) error {
// d.URL = d.URL + "/domains/" + d.Host
// d.Method = "PATCH"
// d.Body = body
// return d.Request
// return nil
// }
