package url

// Domain makes domainEndpoint public
// type Domain interface {
// 	Contacts() string
// 	Privacy() Privacy
// }

// Domain makes domainEndpoint public
type Domain struct {
	url string
}

// Contacts builds out the contacts piece of the URL
func (d Domain) Contacts() string {
	return d.url + "/contacts"
}

// Privacy builds out the privacy piece of the URL
func (d Domain) Privacy() Privacy {
	return Privacy{url: d.url + "/privacy"}
}
