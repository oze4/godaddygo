package client

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
