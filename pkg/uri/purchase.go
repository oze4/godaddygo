package uri

// Purchase knows how to create URL's for 
// the "purchase endpoint"
type Purchase interface {
    String() string
}

type purchase struct {
    *cache
}

func (p *purchase) String() string {
    return p.path + "/domains/purchase"
}