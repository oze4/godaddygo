package core

// API specifies the production base URL for the GoDaddy API (https://api.godaddy.com)
// type API interface {
// 	V1() V1
// }

// Global scoped for now
var (
	apiKey    string
	apiSecret string
)

// APIInterface specifies the production base URL for the GoDaddy API (https://api.godaddy.com)
type APIInterface interface {
	V1Getter
}

// api implements APIInterface
type api struct {
	url string
}

// V1 returns the V1 section of the GoDaddy API
func (a *api) V1() V1Interface {
	return &v1{url: a.url + "/v1"}
}

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func NewProduction(key, secret string) APIInterface {
	apiKey = key
	apiSecret = secret
	return &api{url: "https://api.godaddy.com"}
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func NewDevelopment() APIInterface {
	panic("The OTE (development) section of this library is under construction!")
	// return &api{url: "https://api.ote-godaddy.com"}
}
