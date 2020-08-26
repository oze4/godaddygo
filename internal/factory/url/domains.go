package url

import (
	"strconv"
	"strings"
)

// Domains makes domainsEndpoint public
// type Domains interface {
// 	Agreements([]string, bool, bool) string
// 	Available(string) string
// 	Contacts() Contacts
// }

// Domains is for non-authenticated public domain queries (is domain available? get domain info, etc..)
type Domains struct {
	url string
}

// Agreements builds the agreements piece of the URL
func (d Domains) Agreements(domains []string, privacyRequested, forTransfer bool) string {
	dl := strings.Join(domains, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	return d.url + "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
}

// Available builds the available piece of the URL
func (d Domains) Available(domain string) string {
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	return d.url + "/available?domain=" + domain + "&checkType=FAST&forTransfer=false"
}

// Contacts builds the contacts piece of the URL
func (d Domains) Contacts() Contacts {
	return Contacts{url: d.url + "/contacts"}
}
