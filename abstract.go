package godaddygo

import (
	"context"
)

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
	Add(ctx context.Context, rec []Record) error
	FindByType(ctx context.Context, t string) ([]Record, error)
	FindByTypeAndName(ctx context.Context, t string, n string) ([]Record, error)
	ReplaceByType(ctx context.Context, t string, rec []Record) error
	ReplaceByTypeAndName(ctx context.Context, t string, n string, rec []Record) error
	Update(ctx context.Context, rec []Record) error
	Delete(ctx context.Context, rec Record) error
}

//
// APIEnv represents which endpoint to target (dev|prod)
//
type APIEnv string

func (e APIEnv) String() string {
	switch e {
	case APIProdEnv:
		return "prod"
	case APIDevEnv:
		return "dev"
	default:
		return ""
	}
}

// IsValid determines whether or not the given api env is valid
func (e APIEnv) IsValid() bool {
	return e.String() != ""
}

//
// APIVersion represents which endpoint version to target (v1|v2)
//
type APIVersion string

func (v APIVersion) String() string {
	switch v {
	case APIVersion1:
		return "v1"
	case APIVersion2:
		return "v2"
	default:
		return ""
	}
}

// IsValid determines whether or not the given api version is valid
func (v APIVersion) IsValid() bool {
	return v.String() != ""
}

//
// APIURL represents which URL to target
//
type APIURL string

func (u APIURL) String() string {
	switch u {
	case prodbaseURL:
		return "https://api.godaddy.com"
	case devbaseURL:
		return "https://api.ote-godaddy.com"
	default:
		return ""
	}
}

// IsValid determines whether or not the given api url is valid
func (u APIURL) IsValid() bool {
	return u.String() != ""
}

//
// RecordType represents a DNS record type
//
type RecordType string

func (r RecordType) String() string {
	switch r {
	case RecordTypeA:
		return "A"
	case RecordTypeAAAA:
		return "AAAA"
	case RecordTypeCNAME:
		return "CNAME"
	case RecordTypeMX:
		return "MX"
	case RecordTypeNS:
		return "NS"
	case RecordTypeSOA:
		return "SOA"
	case RecordTypeSRV:
		return "SRV"
	case RecordTypeTXT:
		return "TXT"
	default:
		return ""
	}
}

// IsDeletable determines if the given record can be deleted or not
func (r RecordType) IsDeletable() bool {
	switch r {
	case RecordTypeA, RecordTypeAAAA, RecordTypeCNAME, RecordTypeMX, RecordTypeSRV, RecordTypeTXT:
		return true
	default:
		return false
	}
}

// IsValid determines whether or not the given dns recorsd type is valid
func (r RecordType) IsValid() bool {
	return r.String() != ""
}
