package godaddygo

// URLBuilder struct
type URLBuilder struct{}

// NewURLBuilder returns a new URLBuilder struct
func NewURLBuilder() URLBuilder {
    return URLBuilder{}
}

// APIV1 return api v1 endpoints
func (u URLBuilder) APIV1() APIV1 {
    return APIV1{URL: "https://api.godaddy.com/v1"}
}

// APIV1 does a thing
type APIV1 struct {
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

// DomainsEndpoint makes domainsEndpoint public
type DomainsEndpoint interface {
    Agreements() string
    Available() string
    Contacts() DomainContacts
}

type domainsEndpoint struct {
    URL string
}

func (d domainsEndpoint) Agreements() string {
    return d.URL + "/agreements"
}

func (d domainsEndpoint) Available() string {
    return d.URL + "/available"
} 

func (d domainsEndpoint) Contacts() DomainContacts {
    return domainContacts{URL: d.URL + "/contacts"}
}

type domainContacts struct{
    URL string
}

func (dc domainContacts) Validate() string {
    return dc.URL + "/validate"
}

// DomainContacts does a thing
type DomainContacts interface{
    Validate() string
}