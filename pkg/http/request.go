package http

// Request holds request data
type Request struct {
	APIKey    string
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
