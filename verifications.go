package godaddygo

// Verifications defines who verified purchases, etc..
type Verifications struct {
	DomainName DomainName `json:"domainName,omitempty"`
	RealName   RealName   `json:"realName,omitempty"`
}