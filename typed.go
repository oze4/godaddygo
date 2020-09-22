package godaddygo

import (
	"time"
)

// Domain knows how to interact with the Domains
// GoDaddy API endpoint
type Domain interface {
	Records() Records
	GetDetails() (DomainDetails, error)
}

// DomainDetails defines the details of a domain
type DomainDetails struct {
	AuthCode               string
	ContactAdmin           Contact
	ContactBilling         Contact
	ContactRegistrant      Contact
	ContactTech            Contact
	CreatedAt              time.Time
	DeletedAt              time.Time
	TransferAwayEligibleAt time.Time
	Domain                 string
	DomainID               int
	ExpirationProtected    bool
	Expires                time.Time
	ExposeWhois            bool
	HoldRegistrar          bool
	Locked                 bool
	NameServers            []string
	Privacy                bool
	RenewAuto              bool
	RenewDeadline          time.Time
	Status                 string
	SubAccountID           string
	TransferProtected      bool
	Verifications          Verifications
}

// Contact defines the details of a contact
type Contact struct {
	AddressMailing AddressMailing
	Email          string
	Fax            string
	JobTitle       string
	NameFirst      string
	NameLast       string
	NameMiddle     string
	Organization   string
	Phone          string
}

// AddressMailing defines a mailing address
type AddressMailing struct {
	Address    string
	Address2   string
	City       string
	Country    string
	PostalCode string
	State      string
}

// Verifications defines who verified purchases, etc..
type Verifications struct {
	DomainName DomainName
	RealName   RealName
}

// RealName defines the real name
type RealName struct {
	Status string
}

// DomainName defines a domain name
type DomainName struct {
	Status string
}

// Records knows how to interact with the Records
// GoDaddy API endpoint
type Records interface {
	List() ([]Record, error)
	FindByType(t string) ([]Record, error)
	FindByTypeAndName(t string, n string) ([]Record, error)
	Update(rec Record) error
	Delete(rec Record) error
}

const (
	// RecordTypeA defines A record
	RecordTypeA = "A"
	// RecordTypeAAAA defines AAAA record
	RecordTypeAAAA = "AAAA"
	// RecordTypeCNAME defines CNAME record
	RecordTypeCNAME = "CNAME"
	// RecordTypeMX defines MX record
	RecordTypeMX = "MX"
	// RecordTypeNS defines NS record
	RecordTypeNS = "NS"
	// RecordTypeSOA defines SOA record
	RecordTypeSOA = "SOA"
	// RecordTypeSRV defines SRV record
	RecordTypeSRV = "SRV"
	// RecordTypeTXT defines TXT record
	RecordTypeTXT = "TXT"
)

// Record defines a DNS record
type Record struct {
	Data     string
	Name     string
	Port     int
	Priority int
	Protocol string
	Service  string
	TTL      int
	Type     string
	Weight   int
}
