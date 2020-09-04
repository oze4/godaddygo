package api

// InitProduction targets GoDaddy's production API (https://api.godaddy.com)
func InitProduction(apiKey, apiSecret string) Gateway {
	return Gateway{
		currentRequest: currentRequest{
			isProduction: true,
			apiKey:       apiKey,
			apiSecret:    apiSecret,
		},
	}
}

// InitDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func InitDevelopment(apiKey, apiSecret string) Gateway {
	return Gateway{
		currentRequest: currentRequest{
			isProduction: false,
			apiKey:       apiKey,
			apiSecret:    apiSecret,
		},
	}
}

type currentRequest struct {
	apiKey       string
	apiSecret    string
	isProduction bool
	domainName   string
	version      string
}
