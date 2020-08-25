package godaddygo

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