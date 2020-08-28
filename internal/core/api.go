package core

// APIInterface specifies the production base URL for the GoDaddy API (https://api.godaddy.com)
type APIInterface interface {
	V1Getter
}

// api implements APIInterface
type api struct {
	*request
}

// V1 returns the V1 section of the GoDaddy API
func (a *api) V1() V1Interface {
	a.url = a.url + "/v1"
	return &v1{a.request}
}

// NewProductionAPI targets GoDaddy's production API (https://api.godaddy.com)
func NewProductionAPI(key, secret string) APIInterface {
	return &api{
		&request{
			apiKey:    key,
			apiSecret: secret,
			url:       "https://api.godaddy.com",
		},
	}
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func NewDevelopment() APIInterface {
	panic("The OTE (development) section of this library is under construction!")
	// return &api{url: "https://api.ote-godaddy.com"}
}
