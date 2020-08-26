package client

type client struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
	APIMode   string `json:"apiMode"` // Choice of production or OTE
}
