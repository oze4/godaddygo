package url

// Contacts builds the contacts piece of the URL
type Contacts interface {
	Validate() string
}

type contacts struct {
	url string
}

func (c contacts) Validate() string {
	return c.url + "/validate"
}
