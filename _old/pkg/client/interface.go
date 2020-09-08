package client

// Interface holds GoDaddy API client data
type Interface interface {
	APIKey() string
	APISecret() string
	IsProduction() bool
}

// NewConnection creates a new session, which gives
// us access to storeand retrieve data specific to a
// "connection" of GoDaddy API endpoints
func NewConnection(godaddygoClient Interface) Gateway {
	return &gateway{
		&session{
			isProduction: godaddygoClient.IsProduction(),
			apiKey:       godaddygoClient.APIKey(),
			apiSecret:    godaddygoClient.APISecret(),
		},
	}
}
