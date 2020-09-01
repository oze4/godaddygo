package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// ContactsGetter allows for easy embedding (?)
type ContactsGetter interface {
	Contacts() Contacts
}

// Contacts builds the contacts piece of the URL
type Contacts interface {
	Validate() *http.Request
}

// contacts implements Contacts
type contacts struct {
	*http.Request
}

// Validate builds the validate piece of the URL
func (c *contacts) Validate() *http.Request {
	c.URL = c.URL + "/contacts/validate"
	return c.Request
}
