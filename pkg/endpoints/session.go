package endpoints

import (
	"github.com/oze4/godaddygo/pkg/client"
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
