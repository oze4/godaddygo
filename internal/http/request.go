package http

// GoDaddyRequest holds info related to GoDaddy API requests
type GoDaddyRequest interface {
	Do() interface{}
}

type goDaddyRequest struct {
	url        string // `json:"url"`
	domainName string // `json:"domainName"`
	method     string // `json:"method"`
}

func (r goDaddyRequest) Do() interface{} {
	//TODO: Add logic here for sending HTTP requests with fasthttp
	return new(interface{})
}

// NewGoDaddyRequest returns a new GoDaddy request
func NewGoDaddyRequest(url, domainName, requestMethod string) GoDaddyRequest {
	if r := ValidateMethod(requestMethod); r != true {
		panic("Unacceptable requestMethod (" + requestMethod + ") supplied!")
	}

	return goDaddyRequest{
		url:        url,
		domainName: domainName,
		method:     requestMethod,
	}
}
