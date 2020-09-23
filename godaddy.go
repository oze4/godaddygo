package godaddygo

// New returns a new Gateway based upon a config
func New(c *Config) Gateway {
	return newGateway(c)
}

// NewProduction targets GoDaddy production API
func NewProduction(key string, secret string) Gateway {
	c := newDefaultConfig(key, secret, APIProdEnv)
	return newGateway(c)
}

// NewDevelopment targets GoDaddy development API
func NewDevelopment(key string, secret string) Gateway {
	c := newDefaultConfig(key, secret, APIDevEnv)
	return newGateway(c)
}
