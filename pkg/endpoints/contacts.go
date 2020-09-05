package endpoints

// ContactsGetter allows for easy embedding (?)
type ContactsGetter interface {
	Contacts() Contacts
}

// Contacts builds the contacts piece of the URL
type Contacts interface {
	Validate() error
}

// contacts implements Contacts
type contacts struct {
	connectionInterface
}

// Validate builds the validate piece of the URL
func (c *contacts) Validate() error {
	// c.URL = c.URL + "/contacts/validate"
	// return c.Request
	return nil
}
