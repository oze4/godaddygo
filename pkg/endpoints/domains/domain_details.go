package domains

import (
	"time"
)

// DomainDetails holds information about a GoDaddy domain.
// This is the response when you `GET` info about a domain.
type DomainDetails struct {
	AuthCode               string        `json:"authCode"`
	ContactAdmin           Contact       `json:"contactAdmin"`
	ContactBilling         Contact       `json:"contactBilling"`
	ContactRegistrant      Contact       `json:"contactRegistrant"`
	ContactTech            Contact       `json:"contactTech"`
	CreatedAt              time.Time     `json:"createdAt"`
	DeletedAt              time.Time     `json:"deletedAt"`
	TransferAwayEligibleAt time.Time     `json:"transferAwayEligibleAt"`
	Domain                 string        `json:"domain"`
	DomainID               int           `json:"domainId"`
	ExpirationProtected    bool          `json:"expirationProtected"`
	Expires                time.Time     `json:"expires"`
	ExposeWhois            bool          `json:"exposeWhois"`
	HoldRegistrar          bool          `json:"holdRegistrar"`
	Locked                 bool          `json:"locked"`
	NameServers            []string      `json:"nameServers"`
	Privacy                bool          `json:"privacy"`
	RenewAuto              bool          `json:"renewAuto"`
	RenewDeadline          time.Time     `json:"renewDeadline"`
	Status                 string        `json:"status"`
	SubAccountID           string        `json:"subaccountId"`
	TransferProtected      bool          `json:"transferProtected"`
	Verifications          Verifications `json:"verifications"`
}
