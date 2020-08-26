package url

// Domain makes domainEndpoint public
type Domain interface {
	Contacts() string
	Privacy() Privacy
}

type domain struct {
	url string
}

func (d domain) Contacts() string {
	return d.url + "/contacts"
}

func (d domain) Privacy() Privacy {
	return privacy{url: d.url + "/privacy"}
}
