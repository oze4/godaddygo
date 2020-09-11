package endpoints

// New creates a new Contacts
func newContacts(s *session, domainname string) Contacts {
	s.domainName = domainname
	return &contacts{s}
}

// ContactsGetter makes embedding easier
type ContactsGetter interface {
	Contacts(domainname string) Contacts
}

// Contacts defines Contacts behavior
type Contacts interface {
	Validate() error
}

// contacts implements Contacts
type contacts struct {
	*session
}

// Validate builds the validate piece of the URL
func (c *contacts) Validate() error {
	// c.URL = c.URL + "/contacts/validate"
	// return c.Request
	return nil
}

// Contact holds contact information
type Contact struct {
	AddressMailing AddressMailing `json:"addressMailing,omitempty"`
	Email          string         `json:"email,omitempty"`
	Fax            string         `json:"fax,omitempty"`
	JobTitle       string         `json:"jobTitle,omitempty"`
	NameFirst      string         `json:"nameFirst,omitempty"`
	NameLast       string         `json:"nameLast,omitempty"`
	NameMiddle     string         `json:"nameMiddle,omitempty"`
	Organization   string         `json:"organization,omitempty"`
	Phone          string         `json:"phone,omitempty"`
}

// AddressMailing holds contact address info
type AddressMailing struct {
	Address    string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State      string `json:"state,omitempty"`
}
