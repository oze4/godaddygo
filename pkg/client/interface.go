package client

// NewClient is an "batteries-included" default client
// that can be used out of the box
func NewClient(apikey, apisecret string, isproduction bool) Interface {
	return &defaultClient{
		apiKey:       apikey,
		apiSecret:    apisecret,
		isProduction: isproduction,
	}
}

// Interface defines how a client should behave
// By satisfying this interface, you can use your
// own client (eg: `endpoints.Connect( <yourClient> )` 
type Interface interface {
	APIKey() string
	APISecret() string
	IsProduction() bool
}
