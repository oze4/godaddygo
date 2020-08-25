package godaddygo

// ClientOptions hold the options for a new client
type ClientOptions struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
	APIMode   string `json:"apiMode"`
}

// NewClientOptions returns a pointer to our client options
func NewClientOptions() *ClientOptions {
    return &ClientOptions{
        APIMode: "production",
    }
}