package endpoints

// newV1 creates a new v1
func newV1(s *session) V1 {
	s.apiVersion = "v1"
	return &v1{s}
}

// V1 targets version 1 of the GoDaddy API
type V1 interface {
	Domain(hostname string) Domain
}

type v1 struct {
	*session
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v1) Domain(hostname string) Domain {
	return newDomain(v.session, hostname)
}
