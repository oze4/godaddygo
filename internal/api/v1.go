package api

// V1Getter does a thing
type V1Getter interface {
	V1() V1Interface
}

// V1Interface does a thing
type V1Interface interface {
	DomainGetter
}

type v1 struct {
	*request
}

// Domain is most likely what you're looking for. It allows you to modify domains you control
func (v *v1) Domain(domainName string) DomainInterface {
	return &domain{name: domainName, url: v.url + "/domains/" + domainName}
}
