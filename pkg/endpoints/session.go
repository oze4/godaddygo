package endpoints

import (
	"github.com/oze4/godaddygo/pkg/client"
	"github.com/oze4/godaddygo/pkg/uri"
)

func newSession(clientInterface client.Interface) *session {
	return &session{Interface: clientInterface}
}

// session defines how a session behaves
type session struct {
	client.Interface
	method     string
	path       string
	domainName string
	apiVersion string
}

// URLBuilder wraps `uri.Builder`
func (s *session) URLBuilder() uri.Version {
	// s.IsProduction *has* to already be set by the time we see it here
	// It does not matter if `s.apiVersion` is empty here or not
	// We let the error the API will produce guide us
	return uri.Builder(s.IsProduction()).Version(s.apiVersion)
}
