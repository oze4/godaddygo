package godaddygo

import (
	"errors"

	"github.com/oze4/godaddygo/pkg/session"
)

// connectionBridge holds internal connection info
//  - Embeds session.Interface
//  - "Bridges" whatever client we are given to how we 
//    track settings internally 
type connectionBridge interface {
	session.Interface
	TargetDomain() string
	SetTargetDomain(n string)
	SetAPIVersion(v string)
	getBaseURL() (string, error)
}

// connection implements connectionBridge
type connection struct {
	apiKey       string
	apiSecret    string
	isProduction bool
	apiVersion   string
	targetDomain string
}

// NewConnection creates a new session, which gives 
// us access to storeand retrieve data specific to a
// "connection" of GoDaddy API endpoints
func NewConnection(sessionInterface session.Interface) Gateway {
	return &gateway{
		&connection{
			isProduction: sessionInterface.IsProduction(),
			apiKey:       sessionInterface.APIKey(),
			apiSecret:    sessionInterface.APISecret(),
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

func (c *connection) SetAPIVersion(v string) {
	c.apiVersion = v
}

// getBaseURL builds the "base" of our URL
// It builds out the host plus version. eg: 
// https://api.godaddy.com/v1 for production version 1
func (c *connection) getBaseURL() (string, error) {
	if c.apiVersion == "" {
		return "", errors.New("API version not present")
	}

	url := "https://api-ote.godaddy.com/" // Development
	if c.isProduction {
		url = "https://api.godaddy.com/" // Production
	}

	return url + c.apiVersion, nil
}
