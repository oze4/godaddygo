package uri

const (
	privacyRoot = "/privacy"
)

// Domain is the domains endpoint
type Domain interface {
	String() string
	Records() Records
	RemovePrivacy() string
	PurchasePrivacy() string
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

func (d *domain) RemovePrivacy() string {
	return d.path + privacyRoot
}

func (d *domain) PurchasePrivacy() string {
	return d.path + privacyRoot + "/purchase"
}
