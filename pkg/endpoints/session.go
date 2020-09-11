package endpoints

/**
 *
 * `session` is essentially private cache that is used for each connection to our `endpoints`
 *
 */

import (
	"github.com/oze4/godaddygo/pkg/client"
	"github.com/oze4/godaddygo/pkg/rest"
)

// session is essentially private cache that is used for each connection to our `endpoints`
func newSession(clientInterface client.Interface) *session {
	return &session{
		Request: &rest.Request{
			APIKey:    clientInterface.APIKey(),
			APISecret: clientInterface.APISecret(),
		},
		isProduction: clientInterface.IsProduction(),
	}
}

// session is essentially private cache that is used for each connection to our `endpoints`
type session struct {
	*rest.Request
	isProduction bool
	domainName   string
	apiVersion   string
}
