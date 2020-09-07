package domains

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
