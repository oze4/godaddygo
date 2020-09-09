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

// URLBase returns the "base" GoDaddy API URL
// By "base" we mean the target API plus API version
// *Production Version 1 would return:
// https://api.godaddy.com/v1
func (s *session) urlBase() string {
	u := "https://api-ote.godaddy.com"
	if s.IsProduction() {
		u = "https://api.godaddy.com"
	}
	return u + "/" + s.apiVersion
}

// URLBasePlus wraps URLBase()
// We append any string you want to the end of the
// base URL
func (s *session) URLBasePlus(extra string) string {
	b := s.urlBase()
	return b + extra
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
