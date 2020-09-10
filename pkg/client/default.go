package client

// Default is an "batteries-included" default client
// that can be used out of the box
func Default(apikey, apisecret string, isproduction bool) Interface {
	return &defaultClient{
		apiKey:       apikey,
		apiSecret:    apisecret,
		isProduction: isproduction,
	}
}

// defaultClient implements Interface
type defaultClient struct {
	apiKey       string
	apiSecret    string
	isProduction bool
}

func (c *defaultClient) APIKey() string {
	return c.apiKey
}

func (c *defaultClient) APISecret() string {
	return c.apiSecret
}

func (c *defaultClient) IsProduction() bool {
	return c.isProduction
}
