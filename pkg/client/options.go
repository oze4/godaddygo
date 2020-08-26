package client

import (
	"fmt"
)

// Options hold the options for a new client
type Options struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
	APIMode   string `json:"apiMode"` // `production` or `ote`
}

var apiModes = map[string]string{
	"production": "production",
	"ote":        "ote",
}

func validateOptions(o Options) {
	v, ok := apiModes[o.APIMode]
	fmt.Println(v, ok)
}

// NewOptions returns a pointer to our client options
func NewOptions(apiKey, apiSecret, apiMode string) *Options {
	return &Options{
		APIKey:    apiKey,
		APISecret: apiSecret,
		APIMode:   apiMode,
	}
}
