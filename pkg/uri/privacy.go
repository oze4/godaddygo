package uri

// PrivacyGetter simplifies embedding
type PrivacyGetter interface {
	Privacy() Privacy
}

// Privacy knows how to purchase/remove privacy
// for a domain
type Privacy interface {
	Purchase() string
	Remove() string
}

type privacy struct {
	*cache
}

func (p *privacy) Purchase() string {
	return p.path + "/privacy/purchase"
}

func (p *privacy) Remove() string {
	return p.path + "/privacy"
}
