package url

// Contacts builds the contacts piece of the URL
type Contacts struct {
	url string
}

// Validate builds the validate piece of the URL
func (c Contacts) Validate() string {
	return c.url + "/validate"
}
