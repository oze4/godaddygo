package godaddygo

// Options hold the options for a new client
type Options interface {
	APIKey() string
	APISecret() string
}

type options struct {
	apiKey    string
	apiSecret string
}

func (o *options) APIKey() string {
	return o.apiKey
}

func (o *options) APISecret() string {
	return o.apiSecret
}

// NewOptions returns a pointer to our client options
func NewOptions(apiKey, apiSecret string) Options {
	return &options{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}