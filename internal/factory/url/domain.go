package url

// DomainEndpoint makes domainEndpoint public
type DomainEndpoint interface {
	Contacts() string
	Privacy() Privacy
}

type domainEndpoint struct {
	url string
}

func (d domainEndpoint) Contacts() string {
	return d.url + "/contacts"
}

func (d domainEndpoint) Privacy() Privacy {
	return privacy{url: d.url + "/privacy"}
}
