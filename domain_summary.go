package godaddygo

import (
    "time"
)

// DomainSummary is what gets returned when listing all of your domains
type DomainSummary struct {
	CreatedAt              time.Time `json:"createdAt,omitempty"`
	Domain                 string    `json:"domain,omitempty"`
	DomainID               int       `json:"domainId,omitempty"`
	ExpirationProtected    bool      `json:"expirationProtected,omitempty"`
	Expires                time.Time `json:"expires,omitempty"`
	ExposeWhois            bool      `json:"exposeWhois,omitempty"`
	HoldRegistrar          bool      `json:"holdRegistrar,omitempty"`
	Locked                 bool      `json:"locked,omitempty"`
	NameServers            []string  `json:"nameServers,omitempty"`
	Privacy                bool      `json:"privacy,omitempty"`
	RenewAuto              bool      `json:"renewAuto,omitempty"`
	Renewable              bool      `json:"renewable,omitempty"`
	Status                 string    `json:"status,omitempty"`
	TransferAwayEligibleAt time.Time `json:"transferAwayEligibleAt,omitempty"`
	TransferProtected      bool      `json:"transferProtected,omitempty"`
}