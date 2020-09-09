package uri

import (
	"strconv"
)

// Builder is the start of generating an API URL
func Builder(isproduction bool) Gateway {
	h := "https://api-ote.godaddy.com"
	if isproduction {
		h = "https://api.godaddy.com"
	}

	return &gateway{&cache{h}}
}

// cache holds the current path we are building
type cache struct {
	path string
}

// Gateway allows you to target specific API versions
type Gateway interface {
    // We do not validate the `Version(v string)` parameter!
    // We expect you to have already validated the
    // version string beforehand
	Version(v string) Version
}

type gateway struct {
	*cache
}

// We do not validate the `v string` parameter!
// We expect you to have already validated the
// version string beforehand
func (g *gateway) Version(v string) Version {
	g.path += "/" + v
	return &version{g.cache}
}

// Version are the API versions
type Version interface {
    Domain(domainName string) Domain
    DomainAvailability(domainName string, forTransfer bool) string
}

type version struct {
	*cache
}

func (v *version) Domain(domainName string) Domain {
	v.path += "/domains/" + domainName
	return &domain{v.cache}
}

func (v *version) DomainAvailability(domainName string, forTransfer bool) string {
    return v.path + "/available?domain=" + domainName + "&checkType=FAST&forTransfer=" + strconv.FormatBool(forTransfer)
}

// Domain is the domains endpoint
type Domain interface {
	String() string
    Records() Records
}

type domain struct {
	*cache
}

func (d *domain) String() string {
	return d.path
}

func (d *domain) Records() Records {
	d.path += "/records"
	return &records{d.cache}
}

// Records is the `/domain/<domain>/records` endpoint
type Records interface {
	String() string
	ByType(rectype string) string
	ByTypeName(rectype, recname string) string
}

type records struct {
	*cache
}

func (r *records) String() string {
	return r.path
}

func (r *records) ByType(rectype string) string {
	return r.path + "/" + rectype
}

func (r *records) ByTypeName(rectype, recname string) string {
	return r.ByType(rectype) + "/" + recname
}
