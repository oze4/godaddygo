package endpoints

// newV1 creates a new v1
func newV1(s *session) V1 {
	s.apiVersion = "v1"
	return &v1{s}
}

// V1 targets version 1 of the GoDaddy API
type V1 interface {
	// Domain knows how to interact with domains you own
	// Domain is used to target a specific domain
	//  - Get DNS record(s)
	//  - Modify DNS record(s)
	//  - Create DNS record(s)
	// etc...
	DomainGetter

	// Domains knows how to interact with domains you may or may not
	// own but want to perform tasks on.
	//  - Check if a domain is available for purchase
	//  - Purchase a domin
	//  - List all of the domains you own
	// etc...
	DomainsGetter
}

type v1 struct {
	*session
}

// Domain knows how to interact with domains you own
// Domain is used to target a specific domain
//  - Get DNS record(s)
//  - Modify DNS record(s)
//  - Create DNS record(s)
// etc...
func (v *v1) Domain(domainname string) Domain {
	return newDomain(v.session, domainname)
}

// Domains knows how to interact with domains you may or may not
// own but want to perform tasks on.
//  - Check if a domain is available for purchase
//  - Purchase a domin
//  - List all of the domains you own
// etc...
func (v *v1) Domains() Domains {
	return newDomains(v.session)
}
