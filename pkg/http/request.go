package http

// Request holds request data
// type Request interface {
// 	APIKey() string
// 	APISecret() string
// 	URL() string
// 	Method() string
// }

// Request holds request data
type Request struct {
	APIKey     string
	APISecret string
	Method    string
	URL       string
	Host      string
	Body      []byte
}

// Do sends the http request
func (r *Request) Do() {
	// Logic to send request goes here

	// Need to verify request method!
}

/*
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

// NewRequest returns a new request object to send raw requests to GoDaddy
func NewRequest() Request {
	return &request{}
}
*/
