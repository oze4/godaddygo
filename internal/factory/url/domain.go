package url

// Domain makes domainEndpoint public
// type Domain interface {
// 	Contacts() string
// 	Privacy() Privacy
// }

// Domain makes domainEndpoint public
type Domain struct {
	GoDaddy
}

// Contacts builds out the contacts piece of the URL
func (d Domain) Contacts() GoDaddy {
	// return d.url + "/contacts"
	return GoDaddy{URL: d.URL, DomainName: d.DomainName}
}

// Privacy builds out the privacy piece of the URL
func (d Domain) Privacy() Privacy {
	// return Privacy{url: d.url + "/privacy"}
	return Privacy{
		GoDaddy{URL: d.URL + "/privacy", DomainName: d.DomainName},
	}
}
