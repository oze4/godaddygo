package api

// V1Getter does a thing
type V1Getter interface {
	V1() V1
}

// V1 does a thing
type V1 interface {
	DomainGetter
}

type v1 struct {
	*request
}

// Domain provides domain related info and tasks for the `domains` GoDaddy API endpoint
func (v *v1) Domain(name string) Domain {
	return &domain{name: name, url: v.url + "/domains/" + name}
}
