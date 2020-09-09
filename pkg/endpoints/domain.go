package endpoints

import (
	"encoding/json"
	"time"

	"github.com/oze4/godaddygo/pkg/rest"
)

// newDomain creates a new domain
func newDomain(s *session, domainName string) Domain {
	s.domainName = domainName
	return &domain{s}
}

// Domain implements Domain [interface]
type Domain interface {
	ContactsGetter
	PrivacyGetter
	RecordsGetter
	Agreements(domains []string, privacyRequested, forTransfer bool) error
	GetDetails() (*DomainDetails, error)
}

type domain struct {
	*session
}

func (d *domain) Records() Records {
	return newRecords(d.session)
}

// Contacts builds out the contacts piece of the URL
func (d *domain) Contacts() Contacts {
	return newContacts(d.session)
}

// Privacy builds out the privacy piece of the URL
func (d *domain) Privacy() Privacy {
	return newPrivacy(d.session)
}

// Agreements builds the agreements piece of the URL
func (d *domain) Agreements(domains []string, privacyRequested, forTransfer bool) error {
	/*
		d.URL = d.URL + "/domains"
		doms := append(domains, d.Host)
		dl := strings.Join(doms, ",")
		p := strconv.FormatBool(privacyRequested)
		f := strconv.FormatBool(forTransfer)
		d.URL = "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
		return d.Request
	*/
	return nil
}

// GetDetails gets info on a domain
func (d *domain) GetDetails() (*DomainDetails, error) {
	req := &rest.Request{
		APIKey:    d.APIKey(),
		APISecret: d.APISecret(),
		URL:       d.URLBasePlus("/" + d.domainName),
		Method:    "GET",
	}

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	var details DomainDetails
	if err := json.Unmarshal(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// DomainDetails holds information about a GoDaddy domain.
// This is the response when you `GET` info about a domain.
type DomainDetails struct {
	AuthCode               string        `json:"authCode,omitempty"`
	ContactAdmin           Contacts      `json:"contactAdmin,omitempty"`
	ContactBilling         Contacts      `json:"contactBilling,omitempty"`
	ContactRegistrant      Contacts      `json:"contactRegistrant,omitempty"`
	ContactTech            Contacts      `json:"contactTech,omitempty"`
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

// Update updates a domain
// func (d *domain) Update(body []byte) error {
// d.URL = d.URL + "/domains/" + d.Host
// d.Method = "PATCH"
// d.Body = body
// return d.Request
// return nil
// }
