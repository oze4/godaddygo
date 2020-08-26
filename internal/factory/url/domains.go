package url

import (
	"strconv"
	"strings"
)

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
