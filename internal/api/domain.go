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

func (d *domain) attach() {
	d.URL = d.URL + "/domains/" + d.Host
}

// Contacts builds out the contacts piece of the URL
func (d *domain) Contacts() Contacts {
	d.attach()
	// d.URL = d.URL + "/contacts"
	return &contacts{d.Request}
}

// Privacy builds out the privacy piece of the URL
func (d *domain) Privacy() Privacy {
	d.attach()
	return &privacy{d.Request}
}

// Agreements builds the agreements piece of the URL
func (d *domain) Agreements(domains []string, privacyRequested, forTransfer bool) *http.Request {
	d.attach()
	dl := strings.Join(domains, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	d.URL = "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
	return d.Request
	// return d.URL + "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
}

// Available builds the available piece of the URL
func (d *domain) Available() *http.Request {
	d.attach()
	d.Method = "GET"
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	d.URL = d.URL + "/available?domain=" + d.Host + "&checkType=FAST&forTransfer=false"
	return d.Request
}

// Records builds the DNS record piece of the URL
func (d *domain) Records() Records {
	d.attach()
	return &records{d.Request}
}

// GetDetails gets info on a domain
func (d *domain) GetDetails() *http.Request {
	d.attach()
	d.Method = "GET"
	return d.Request
}

// Delete deletes a domain
func (d *domain) Delete() {
	d.attach()
	d.Method = "DELETE"
	//TODO: Delete logic here
}

// Update updates a domain
func (d *domain) Update() {
	d.attach()
	d.Method = "PATCH"
	//TODO: Update logic here
}
