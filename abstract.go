package godaddygo

import (
	"context"
	"net/http"
	"time"
)

/** ---------
 * constants
 --------- */

const (
	/** env vars, used for connection parameters **/

	// APIProdEnv targets the production Gateway
	APIProdEnv = "prod"
	// APIDevEnv targets the development Gateway
	APIDevEnv = "dev"

	/** api versions **/

	// APIVersion1 specifies version 1
	APIVersion1 = "v1"
	// APIVersion2 specifies version 2
	APIVersion2 = "v2"

	/** dns record types **/

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

	/** godaddy api url's **/

	prodbaseURL = "https://api.godaddy.com"
	devbaseURL  = "https://api.ote-godaddy.com"

	/** http methods so we don't have to import http everywhere
	  only doing this for common (or used) methods **/

	// MethodGet is global shortcut for http.MethodGet
	MethodGet = http.MethodGet
	// MethodPost is global shortcut for http.MethodPost
	MethodPost = http.MethodPost
	// MethodPut is global shortcut for http.MethodPut
	MethodPut = http.MethodPut
	// MethodPatch is global shortcut for http.MethodPatch
	MethodPatch = http.MethodPatch
	// MethodDelete is global shortcut for http.MethodDelete
	MethodDelete = http.MethodDelete
)

/** -----------
 * interfaces
 ----------- */

// API knows which version to target
type API interface {
	V1() V1
	V2() V2
}

// V1 knows how to interact with GoDaddy Gateway version 1
type V1 interface {
	Domain(name string) Domain
	ListDomains(ctx context.Context) ([]DomainSummary, error)
	CheckAvailability(ctx context.Context, name string, forTransfer bool) (DomainAvailability, error)
	PurchaseDomain(ctx context.Context, dom DomainDetails) error
}

// V2 knows how to interact with GoDaddy Gateway version 2
type V2 interface{}

// Domain knows how to interact with the Domains GoDaddy Gateway endpoint
type Domain interface {
	Records() Records
	GetDetails(ctx context.Context) (*DomainDetails, error)
}

// Records knows how to interact with the Records GoDaddy Gateway endpoint
type Records interface {
	List(ctx context.Context) ([]Record, error)
	FindByType(ctx context.Context, t string) ([]Record, error)
	FindByTypeAndName(ctx context.Context, t string, n string) ([]Record, error)
	Update(ctx context.Context, rec Record) error
	Delete(ctx context.Context, rec Record) error
}

/** --------
 * structs
 -------- */

// Config holds connection options
// use NewConfig to create a new config
type Config struct {
	client     *http.Client
	key        string // key is the api key
	secret     string // secret is the api secret
	baseURL    string // we take care of this
	env        string // env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
	version    string // version should be `v1` or `v2`, use APIVersion1 or APIVersion2
	domainName string // dns zone to target

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
	Data     string `json:"data,omitempty"`
	Name     string `json:"name,omitempty"`
	Port     int    `json:"port,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Service  string `json:"service,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	Type     string `json:"type,omitempty"`
	Weight   int    `json:"weight,omitempty"`
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
