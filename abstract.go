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
