package client

// This is how you can build your own client

// NewDefaultClient creates a new default client
func NewDefaultClient(prod bool, key, secret string) Interface {
	return &defaultClient{
		isProd:    prod,
		apiKey:    key,
		apiSecret: secret,
	}
}

// defaultClient implements Interface (aka client.Interface)
type defaultClient struct {
	isProd    bool
	apiKey    string
	apiSecret string
}

// IsProduction determiens which base URL to use
func (c *defaultClient) IsProduction() bool {
	return c.isProd
}

// APIKey holds the api key
func (c *defaultClient) APIKey() string {
	return c.apiKey
}

// APISecret holds the API secret
func (c *defaultClient) APISecret() string {
	return c.apiSecret
}
