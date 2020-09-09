package endpoints

import (
	"github.com/oze4/godaddygo/pkg/client"
	uri "github.com/oze4/godaddygo/pkg/uri"
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
// We assume that the session apiVersion and
// session IsProduction() values are set prior
// to this
// These should both be set before needing to construct
// a URL, which is why we return `uri.Version` and not `uri.Gateway`
// this is for convenience
func (s *session) URLBuilder() uri.Version {
	return uri.Builder(s.IsProduction()).Version(s.apiVersion)
}
