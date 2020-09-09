package endpoints

import (
	"github.com/oze4/godaddygo/pkg/client"
)

// Connect allows you to get a new gateway
func Connect(c client.Interface) Gateway {
	s := newSession(c)
	return newGateway(s)
}

func newGateway(s *session) Gateway {
	return &gateway{s}
}

// Gateway knows which GoDaddy API version to target
type Gateway interface {
	V1() V1
	V2() V2
}

type gateway struct {
	*session
}

func (g *gateway) V1() V1 {
	return newV1(g.session)
}

func (g *gateway) V2() V2 {
	return newV2(g.session)
}

/**
 * Convenience functions
 */

// ConnectProduction is an convenience function
// returns a new Gateway which targets GoDaddy
// production API
func ConnectProduction(apikey, apisecret string) Gateway {
	isProduction := true
	c := client.Default(apikey, apisecret, isProduction)
	s := newSession(c)
	return newGateway(s)
}

// ConnectDevelopment is an convenience function
// returns a new Gateway which targets GoDaddy
// development API
func ConnectDevelopment(apikey, apisecret string) Gateway {
	isProduction := false
	c := client.Default(apikey, apisecret, isProduction)
	s := newSession(c)
	return newGateway(s)
}
