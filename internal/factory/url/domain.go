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
	return GoDaddy{url: d.url, domainName: d.domainName}
}

// Privacy builds out the privacy piece of the URL
func (d Domain) Privacy() Privacy {
	// return Privacy{url: d.url + "/privacy"}
	return Privacy{
		GoDaddy{url: d.url + "/privacy", domainName: d.domainName},
	}
}

func a() {
	u := NewBuilder().Production().V1().Domain("ostrike.com")
	t(u.GoDaddy)
}

func t(g GoDaddy) {

}
