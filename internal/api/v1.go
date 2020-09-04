package api

// V1 implements V1
type v1 struct {
	currentRequest
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v v1) Domain(hostname string) Domain {
	return &domain{domainName: hostname}
}
