package godaddygo

import (
	"net/http"
	"time"
)

// This is a comment so my IDE quits complaining to me
const (
	// Env variables - which godaddy api to target
	APIProdEnv APIEnv = "prod"
	APIDevEnv         = "dev"

	// URLs representation of env variable
	prodbaseURL APIURL = "https://api.godaddy.com"
	devbaseURL         = "https://api.ote-godaddy.com"

	// Allowed API Versions - which version of the godaddy api to target
	APIVersion1 APIVersion = "v1"
	APIVersion2            = "v2"

	// DNS record types
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA             = "AAAA"
	RecordTypeCNAME            = "CNAME"
	RecordTypeMX               = "MX"
	RecordTypeNS               = "NS"
	RecordTypeSOA              = "SOA"
	RecordTypeSRV              = "SRV"
	RecordTypeTXT              = "TXT"

	// Domain statuses (added "Domain" prefix to legacy constants)
	DomainStatusActive    = "ACTIVE"
	DomainStatusCancelled = "CANCELLED"
	// Legacy Domain statuses (to support rename)
	StatusActive    = "ACTIVE"
	StatusCancelled = "CANCELLED"
)

// Config holds connection options. Use `.NewConfig` to create a new config
type Config struct {
	client     *http.Client
	key        string     // key is the api key
	secret     string     // secret is the api secret
	baseURL    APIURL     // we take care of this
	env        APIEnv     // env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
	version    APIVersion // version should be `v1` or `v2`, use APIVersion1 or APIVersion2
	domainName string     // dns zone to target
}

// DomainDetails defines the details of a domain
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
	SubAccountID           string        `json:"subAccountId,omitempty"`
	TransferProtected      bool          `json:"transferProtected,omitempty"`
	Verifications          Verifications `json:"verifications,omitempty"`
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
	Address    string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State      string `json:"state,omitempty"`
}

// FullAddress returns the full address (Address + Address2)
func (am *AddressMailing) FullAddress() string {
	return am.Address + " " + am.Address2
}

// Verifications defines who verified purchases, etc..
type Verifications struct {
	DomainName DomainName `json:"domainName,omitempty"`
	RealName   RealName   `json:"realName,omitempty"`
}

// RealName defines the real name
type RealName struct {
	Status string `json:"status,omitempty"`
}

// DomainName defines a domain name
type DomainName struct {
	Status string
}

// Record defines a DNS record
type Record struct {
	Data     string     `json:"data,omitempty"`
	Name     string     `json:"name,omitempty"`
	Port     int        `json:"port,omitempty"`
	Priority int        `json:"priority,omitempty"`
	Protocol string     `json:"protocol,omitempty"`
	Service  string     `json:"service,omitempty"`
	TTL      int        `json:"ttl,omitempty"`
	Type     RecordType `json:"type,omitempty"`
	Weight   int        `json:"weight,omitempty"`
}

// DomainSummary is what gets returned when listing all of your domains
type DomainSummary struct {
	CreatedAt              time.Time
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
	Renewable              bool
	Status                 string
	TransferAwayEligibleAt time.Time
	TransferProtected      bool
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
