package domains

// Verifications holds verification info about a domain.
type Verifications struct {
	DomainName DomainName `json:"domainName,omitempty"`
	RealName   RealName   `json:"realName,omitempty"`
}

// DomainName holds info about domain name
type DomainName struct {
	Status string `json:"status,omitempty"`
}

// RealName holds real name info about a domain.
type RealName struct {
	Status string `json:"status,omitempty"`
}
