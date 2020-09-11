package endpoints

import (
	"encoding/json"
	"time"
)

func newDomain(s *session, n string) Domain {
	s.domainName = n
	return &domain{s}
}

// DomainGetter simplifies embedding
type DomainGetter interface {
	Domain(domainname string) Domain
}

// Domain knows how to interact with domains you own
// Domain is used to target a specific domain
//  - Get DNS record(s)
//  - Modify DNS record(s)
//  - Create DNS record(s)
// etc...
type Domain interface {
	ContactsGetter
	PrivacyGetter
	RecordsGetter
	GetDetails() (*DomainDetails, error)
}

// domain implements Domain
type domain struct {
	*session
}

// Records provides access to the 'records endpoint'
func (d *domain) Records() Records {
	return newRecords(d.session)
}

// Contacts provides access to the 'contacts endpoint'
func (d *domain) Contacts() Contacts {
	return newContacts(d.session)
}

// Privacy provides access to the 'privacy endpoint'
func (d *domain) Privacy() Privacy {
	return newPrivacy(d.session)
}

// GetDetails gets details for the specified domain
// returns `*DomainDetails`
func (d *domain) GetDetails() (*DomainDetails, error) {
	d.Method = "GET"
	d.URL = d.URLBuilder().Domain(d.domainName).String()

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

/*
  {
    "createdAt": "2015-06-15T13:10:43.000Z",
    "domain": "000.biz",
    "domainId": 1002111,
    "expirationProtected": false,
    "expires": "2016-06-14T23:59:59.000Z",
    "exposeWhois": false,
    "holdRegistrar": false,
    "locked": true,
    "nameServers": null,
    "privacy": false,
    "renewAuto": true,
    "renewable": false,
    "status": "TRANSFERRED_OUT",
    "transferAwayEligibleAt": "2016-07-29T23:59:59.000Z",
    "transferProtected": false
  }
*/

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
