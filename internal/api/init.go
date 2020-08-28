package api

// NewProductionAPI targets GoDaddy's production API (https://api.godaddy.com)
func NewProductionAPI(key, secret string) Gateway {
	return &gateway{
		&request{
			apiKey:    key,
			apiSecret: secret,
			url:       "https://api.godaddy.com",
		},
	}
}

// NewDevelopmentAPI targets GoDaddy's development API (https://api.ote-godaddy.com)
func NewDevelopmentAPI() Gateway {
	panic("The OTE (development) section of this library is under construction!")
	// return &api{url: "https://api.ote-godaddy.com"}
}