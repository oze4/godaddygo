package domain

import (
	"github.com/oze4/godaddygo/pkg/rest"
)

// New creates a new Contacts
func newContacts(c *rest.Config) ContactsInterface {
	return &contacts{c}
}

// ContactsGetter makes embedding easier
type ContactsGetter interface {
	ContactsInterface() ContactsInterface
}

// ContactsInterface defines Contacts behavior
type ContactsInterface interface {
	Validate() error
}

// contacts implements Contacts
type contacts struct {
	*rest.Config
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
