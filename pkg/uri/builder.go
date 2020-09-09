package uri

// Builder is the start of generating an API URL
func Builder(isproduction bool) Gateway {
	h := "https://api-ote.godaddy.com"
	if isproduction {
		h = "https://api.godaddy.com"
	}

	return &gateway{&cache{h}}
}

// cache holds the current path we are building
type cache struct {
	path string
}

// Gateway allows you to target specific API versions
type Gateway interface {
	Versions(v string) Versions
}

type gateway struct {
	*cache
}

func (g *gateway) Versions(v string) Versions {
	g.path += "/" + v
	return &versions{g.cache}
}

// Versions are the API versions
type Versions interface {
	Domain(domainName string) Domain
}

type versions struct {
	*cache
}

func (v *versions) Domain(domainName string) Domain {
	v.path += "/domains/" + domainName
	return &domain{v.cache}
}

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

// Records is the `/domain/<domain>/records` endpoint
type Records interface {
	String() string
	GetByType(rectype string) string
	GetByTypeName(rectype, recname string) string
}

type records struct {
	*cache
}

func (r *records) String() string {
	return r.path
}

func (r *records) GetByType(rectype string) string {
	return r.path + "/" + rectype
}

func (r *records) GetByTypeName(rectype, recname string) string {
	return r.GetByType(rectype) + "/" + recname
}
