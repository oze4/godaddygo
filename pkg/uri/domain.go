package uri

// Domain is the domains endpoint
type Domain interface {
	String() string
	Records() Records
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