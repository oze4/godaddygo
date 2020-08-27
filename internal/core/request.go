package core

// Request implements RequestInterface
type Request interface {
	APIKey() string
	APISecret() string
	URL() string
	Method() string
}

type request struct {
	apiKey    string
	apiSecret string
	url       string
	method    string
}

func (r *request) APIKey() string {
	return r.apiKey
}

func (r *request) APISecret() string {
	return r.apiSecret
}

func (r *request) URL() string {
	return r.url
}

func (r *request) Method() string {
	return r.method
}
