package godaddygo

import (
    "strings"
    "strconv"
)

// URLBuilder struct
type URLBuilder struct{}

// NewURLBuilder returns a new URLBuilder struct
func NewURLBuilder() URLBuilder {
    return URLBuilder{}
}

// APIV1 return api v1 endpoints
func (u URLBuilder) APIV1() APIV1 {
    return APIV1{url: "https://api.godaddy.com/v1"}
}

// APIV1 does a thing
type APIV1 struct {
    url string
}

// Domains is meant for general stuff
// Like searching for available domains, or
// viewing public info on existing domains.
func (a APIV1) Domains() DomainsEndpoint {
    e := domainsEndpoint{url: a.url + "/domains"}
    return e;
}

// Domain is meant to target a specific domain you own
func (a APIV1) Domain(d string) DomainEndpoint {
    e := domainEndpoint{url: a.url + "/domains/" + d}
    return e;
}

// DomainEndpoint makes domainEndpoint public
type DomainEndpoint interface {
    Contacts() string
    Privacy() Privacy
}

type domainEndpoint struct {
    url string
}

func (d domainEndpoint) Contacts() string {
    return d.url + "/contacts"
}

func (d domainEndpoint) Privacy() Privacy {
    return privacy{url: d.url + "/privacy"}
}

// Privacy is a thing
type Privacy interface{}

type privacy struct {
    url string
}

// DomainsEndpoint makes domainsEndpoint public
type DomainsEndpoint interface {
    Agreements(domains []string, privacyRequested, forTransfer bool) string
    Available(domain string) string
    Contacts() Contacts
}

type domainsEndpoint struct {
    url string
}

func (d domainsEndpoint) Agreements(domains []string, privacyRequested, forTransfer bool) string {
    dl := strings.Join(domains, ",")
    p := strconv.FormatBool(privacyRequested)
    f := strconv.FormatBool(forTransfer)
    return d.url + "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
}

func (d domainsEndpoint) Available(domain string) string {
    //TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
    return d.url + "/available?domain=" + domain + "&checkType=FAST&forTransfer=false"
} 

func (d domainsEndpoint) Contacts() Contacts {
    return contacts{url: d.url + "/contacts"}
}

// Contacts builds the contacts piece of the URL
type Contacts interface{
    Validate() string
}

type contacts struct{
    url string
}

func (c contacts) Validate() string {
    return c.url + "/validate"
}