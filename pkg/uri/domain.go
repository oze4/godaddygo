package uri

// Domain is the domains endpoint
type Domain interface {
	String() string
	Records() Records
	PrivacyGetter
}

type domain struct {
	*cache
}

func (d *domain) String() string {
	return d.path
}

func (d *domain) Records() Records {
	d.path += "/records"
	return &records{d.cache}
}

func (d *domain) Privacy() Privacy {
	return &privacy{d.cache}
}
