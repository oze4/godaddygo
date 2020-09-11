package uri

import (
	"strconv"
)

// Version are the API versions
type Version interface {
	Domain(domainName string) Domain
	DomainAvailability(domainName string, forTransfer bool) string
	PurchaseDomain() string
}

type version struct {
	*cache
}

func (v *version) Domain(domainName string) Domain {
	v.path += "/domains/" + domainName
	return &domain{v.cache}
}

func (v *version) DomainAvailability(domainName string, forTransfer bool) string {
	return v.path + "/domains/available?domain=" + domainName + "&checkType=FAST&forTransfer=" + strconv.FormatBool(forTransfer)
}

func (v *version) PurchaseDomain() string {
	return v.path + "/domains/purchase"
}