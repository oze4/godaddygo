package core

// V1Getter does a thing
type V1Getter interface {
	V1() V1Interface
}

// V1Interface does a thing
type V1Interface interface {
	Domain(domainName string) Domain
}

type v1 struct {
	url string
}

// Domain is most likely what you're looking for. It allows you to modify domains you control
func (v *v1) Domain(domainName string) Domain {
	return &domain{name: domainName, url: v.url + "/domains/" + domainName}
}
