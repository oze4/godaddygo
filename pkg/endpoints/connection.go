package endpoints

import (
	"errors"

	"github.com/oze4/godaddygo/pkg/session"
)

// connectionInterface holds internal connection info
// embeds session.Interface
type connectionInterface interface {
	session.Interface
	constructRequestURLWithVersion() (string, error)
	TargetDomain() string
	SetTargetDomain(n string)
}

// connection implements connectionInterface
// (Personal Note: ultimately, we want to implement session.Interface)
//TODO delete personal note sometime
type connection struct {
	apiKey       string
	apiSecret    string
	isProduction bool
	apiVersion   string
	targetDomain string
}

// NewConnection creates a new session
func NewConnection(s session.Interface) Gateway {
	return &gateway{
		&connection{
			isProduction: s.IsProduction(),
			apiKey:       s.APIKey(),
			apiSecret:    s.APISecret(),
		},
	}
}

func (c *connection) APIKey() string {
	return c.apiKey
}

func (c *connection) APISecret() string {
	return c.apiSecret
}

func (c *connection) IsProduction() bool {
	return c.isProduction
}

func (c *connection) APIVersion() string {
	return c.apiVersion
}

func (c *connection) TargetDomain() string {
	return c.targetDomain
}

func (c *connection) SetTargetDomain(n string) {
	c.targetDomain = n
}

// ConstructRequestURLWithVersion builds the "base" of our URL
// it builds out the host plus version: https://api.godaddy.com/v1
// for production version 1, for example
func (c *connection) constructRequestURLWithVersion() (string, error) {
	if c.apiVersion == "" {
		return "", errors.New("API version not present")
	}

	url := "https://api-ote.godaddy.com/" // Development
	if c.isProduction {
		url = "https://api.godaddy.com/" // Production
	}

	return url + c.apiVersion, nil
}
