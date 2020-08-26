package client

// Options hold the options for a new client
type Options struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
	APIMode   string `json:"apiMode"`
}

// NewOptions returns a pointer to our client options
func NewOptions() *Options {
	return &Options{
		APIMode: "production",
	}
}