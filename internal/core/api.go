package core

// API specifies the production base URL for the GoDaddy API (https://api.godaddy.com)
type API interface {
	V1() V1
}

type api struct {
	url string
}

// V1 returns the V1 section of the GoDaddy API
func (a api) V1() V1 {
	return v1{url: a.url + "/v1"}
}

// NewProduction targets GoDaddy's production API (https://api.godaddy.com)
func NewProduction() API {
	return api{url: "https://api.godaddy.com"}
}

// NewDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func NewDevelopment() API {
	panic("The OTE (development) section of this library is under construction!")
	// return api{url: "https://api.ote-godaddy.com"}
}
