package api

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func NewProduction(key, secret string) Gateway {
	return &gateway{
		&request{
			apiKey:    key,
			apiSecret: secret,
			url:       "https://api.godaddy.com",
		},
	}
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func NewDevelopment() Gateway {
	panic("The OTE (development) section of this library is under construction!")
	// return &api{url: "https://api.ote-godaddy.com"}
}
