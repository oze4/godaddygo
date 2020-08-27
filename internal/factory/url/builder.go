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
	return Production{url: "https://api.godaddy.com"}
}

// OTE returns development API endpoints
func (b Builder) OTE() OTE {
	panic("The OTE section of this library is under construction!")
	// return OTE{url: "https://api.ote-godaddy.com"}
}
