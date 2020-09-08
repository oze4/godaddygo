package versions

// V2 targets version 1 of the GoDaddy API
type V2 interface {
    // v2 endpoints here
    // mirror v1
}

type v2 struct {
	*rest.Config
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v2) Domain(hostname string) domain.Interface {
	v.setTargetDomain(hostname)
	return &domain{v.meta}
}