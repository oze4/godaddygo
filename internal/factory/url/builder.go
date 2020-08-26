package url

// Builder does a thing
type Builder struct {
	url string
}

// NewBuilder returns a new URLBuilder struct
func NewBuilder() Builder {
	return Builder{}
}

// Production return production API endpoints
func (b Builder) Production() Production {
	return Production{
		GoDaddy{URL: "https://api.godaddy.com"},
	}
}

// OTE returns development API endpoints
func (b Builder) OTE() OTE {
	return OTE{
		GoDaddy{URL: "https://api.ote-godaddy.com"},
	}
}

// Production specifies the production base URL for the GoDaddy API (https://api.godaddy.com)
type Production struct {
	GoDaddy
}

// V1 returns the V1 section of the GoDaddy API
func (p Production) V1() V1 {
	return V1 {
		GoDaddy{URL: p.URL + "/v1"},
	}
}

// OTE specifies the OTE (development) base URL for the GoDaddy API (https://api.ote-godaddy.com)
type OTE struct {
	GoDaddy
}

// UnderConstruction informs users this section is not ready
func (o OTE) UnderConstruction() {
	panic("The OTE section of this library is under construction!")
}
