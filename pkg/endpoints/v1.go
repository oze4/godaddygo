package endpoints

// newV1 creates a new v1
func newV1(s *session) V1 {
	s.apiVersion = "v1"
	return &v1{s}
}

// V1 targets version 1 of the GoDaddy API
type V1 interface {
	DomainsGetter
}

type v1 struct {
	*session
}

// Domains knows how to interact with the domains endpoint
func (v *v1) Domains() Domains {
	return newDomains(v.session)
}
