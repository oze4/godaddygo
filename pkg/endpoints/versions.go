package endpoints

// V1 targets version 1 of the GoDaddy API
type V1 interface {
	Domain(hostname string) Domain
}

type v1 struct {
	connectionInterface
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v v1) Domain(hostname string) Domain {
	v.SetTargetDomain(hostname)
	return &domain{v.connectionInterface}
}
