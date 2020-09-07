package rest

import (
	"strings"
)

// Config holds configuration related data for REST requests
type Config struct {
	// GoDaddy API Key, note that the prod and dev API's have unique API keys/secrets
	APIKey string
	// GoDaddy API Secret, note that the prod and dev API's have unique API keys/secrets
	APISecret string
	// APIVersion holds determines which GoDaddy API to target
	APIVersion string
	// IsProduction determines which URL to send our request to
	IsProduction bool
	// HTTP REST method we validate this
	Method string
	// The path you wish to send your request to
	// Everything after the host (ex of host: https://google.com/) is considered the path
	Path string
	// An example of Host would be: https://google.com Notice the Host includes protocol
	Host string
	// The body of your request, if you need one
	Body []byte
}

// URL returns a properly formatted URL
func (c *Config) URL() string {
	h := "https://api-ote.godaddy.com/" // Development
	if c.IsProduction {
		h = "https://api.godaddy.com/"
	}
	return h + c.Path
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
