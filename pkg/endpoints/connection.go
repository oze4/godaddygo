package endpoints

import (
	"errors"
	"github.com/oze4/godaddygo/pkg/session"
)

// con (short for connection) holds internal connection info
type currentConnection interface {
	session.Interface
	constructRequestURLWithVersion() (string, error)
}

// connection implements currentConnection
type connection struct {
	apiKey       string
	apiSecret    string
	isProduction bool
	apiVersion   string
	targetDomain string
}

// NewConnection creates a new session
func NewConnection(isproduction bool, apikey, apisecret, apiversion string) session.Interface {
	return &connection{
		apiKey:       apikey,
		apiSecret:    apisecret,
		isProduction: isproduction,
		apiVersion:   apiversion,
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
