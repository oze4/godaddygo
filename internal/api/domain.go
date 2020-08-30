package api

import (
	"strconv"
	"strings"

	"github.com/oze4/godaddygo/pkg/http"
)

// DomainGetter returns
type DomainGetter interface {
	Domain(string) Domain
}

// Domain represents the `domains` GoDaddy API endpoint
type Domain interface {
	ContactsGetter
	PrivacyGetter
	RecordsGetter
	Agreements([]string, bool, bool) *http.Request
	Available() *http.Request
	GetDetails() *http.Request
}

// domain implements Domain [interface]
type domain struct {
	*http.Request
}

func (d *domain) attach(attachDomainName bool) {
	if attachDomainName {
		d.URL = d.URL + "/domains/" + d.Host
		return
	}
	d.URL = d.URL + "/domains"
}

// Contacts builds out the contacts piece of the URL
func (d *domain) Contacts() Contacts {
	d.attach(true)
	return &contacts{d.Request}
}

// Privacy builds out the privacy piece of the URL
func (d *domain) Privacy() Privacy {
	d.attach(true)
	return &privacy{d.Request}
}

// Agreements builds the agreements piece of the URL
func (d *domain) Agreements(domains []string, privacyRequested, forTransfer bool) *http.Request {
	d.attach(false)
	dl := strings.Join(domains, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	d.URL = "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
	return d.Request
}

// Available builds the available piece of the URL
func (d *domain) Available() *http.Request {
	d.attach(false)
	d.Method = "GET"
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	d.URL = d.URL + "/available?domain=" + d.Host + "&checkType=FAST&forTransfer=false"
	return d.Request
}

// Records builds the DNS record piece of the URL
func (d *domain) Records() Records {
	d.attach(true)
	return &records{d.Request}
}

// GetDetails gets info on a domain
func (d *domain) GetDetails() *http.Request {
	d.attach(true)
	d.Method = "GET"
	return d.Request
}

// Delete deletes a domain
func (d *domain) Delete() *http.Request {
	d.attach(true)
	d.Method = "DELETE"
	return d.Request
}

// Update updates a domain
func (d *domain) Update(body []byte) *http.Request {
	d.attach(true)
	d.Method = "PATCH"
	d.Body = body
	return d.Request
}
