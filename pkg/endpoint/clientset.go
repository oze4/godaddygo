package endpoint

import (
	domainsEndpoint "github.com/oze4/godaddygo/pkg/endpoint/domains"
)

// ClientSet allows you to talk to different endpoints
type ClientSet interface {
	Domains(name string) domainsEndpoint.Domain
}
