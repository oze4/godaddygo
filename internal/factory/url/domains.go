package url

import (
	"strconv"
	"strings"
)

// Domains makes domainsEndpoint public
type Domains interface {
	Agreements([]string, bool, bool) string
	Available(string) string
	Contacts() Contacts
}

type domains struct {
	url string
}

func (d domains) Agreements(domains []string, privacyRequested, forTransfer bool) string {
	dl := strings.Join(domains, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	return d.url + "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
}

func (d domains) Available(domain string) string {
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	return d.url + "/available?domain=" + domain + "&checkType=FAST&forTransfer=false"
}

func (d domains) Contacts() Contacts {
	return contacts{url: d.url + "/contacts"}
}
