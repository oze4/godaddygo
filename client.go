package godaddygo

import (
	"github.com/oze4/godaddygo/pkg/endpoints"
)

// ConnectProduction returns the production endpoints Gateway
func ConnectProduction(apikey, apisecret string) endpoints.Gateway {
	return endpoints.ConnectProduction(apikey, apisecret)
}

// ConnectDevelopment returns the development endpoints Gateway
func ConnectDevelopment(apikey, apisecret string) endpoints.Gateway {
	return endpoints.ConnectDevelopment(apikey, apisecret)
}
