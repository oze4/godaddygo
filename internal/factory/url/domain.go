package url

import (
	"strconv"
	"strings"
)

// Domain makes domainEndpoint public
type Domain struct {
	url  string
	name string
}

// Contacts builds out the contacts piece of the URL
func (d Domain) Contacts() Contacts {
	// return d.url + "/contacts"
	return Contacts{url: d.url}
}

// Privacy builds out the privacy piece of the URL
func (d Domain) Privacy() Privacy {
	// return Privacy{url: d.url + "/privacy"}
	return Privacy{url: d.url + "/privacy"}
}

// Agreements builds the agreements piece of the URL
func (d Domain) Agreements(domains []string, privacyRequested, forTransfer bool) string {
	dl := strings.Join(domains, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	return d.url + "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
}

// Available builds the available piece of the URL
func (d Domain) Available() string {
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	return d.url + "/available?domain=" + d.name + "&checkType=FAST&forTransfer=false"
}

// Records builds the DNS record piece of the URL
func (d Domain) Records() Records {
	return Records{domain: d}
}

// Get gets info on a domain
func (d Domain) Get() {
	//TODO: Get logic here
}

// Delete deletes a domain
func (d Domain) Delete() {
	//TODO: Delete logic here
}

// Update updates a domain
func (d Domain) Update() {
	//TODO: Update logic here
}
