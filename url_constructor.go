package godaddygo

// URLBuilder struct
type URLBuilder struct{}

// APIV1 does a thing
type APIV1 struct {
    URL string
}

type domainsEndpoint struct {
    URL string
}

// DomainsEndpoint makes domainsEndpoint public
type DomainsEndpoint interface {}

type stepThree struct {
    URL string
}

// NewURLBuilder returns a new URLBuilder struct
func NewURLBuilder() URLBuilder {
    return URLBuilder{}
}

// APIV1 return api v1 endpoints
func (u URLBuilder) APIV1() APIV1 {
    return APIV1{URL: "https://api.godaddy.com/v1"}
}

// Domains does a thing
func (a APIV1) Domains() DomainsEndpoint {
    e := domainsEndpoint{URL: a.URL + "/domains"}
    return e;
}

// Domain does a thing
func (a APIV1) Domain(d string) DomainsEndpoint {
    e := domainsEndpoint{URL: a.URL + "/domains/" + d}
    return e;
}

func (d domainsEndpoint) Agreements() string {
    return d.URL + "/agreements"
}

// DomainList does a list thing
func (d domainsEndpoint) DomainList() stepThree {
    st := stepThree{URL: d.URL + "/domains"}
    return st
}