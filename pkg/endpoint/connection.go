package endpoint

/*
import (
	"errors"
)

// session implements meta
type meta struct {
	apiKey       string
	apiSecret    string
	isProduction bool
	apiVersion   string
	domain       string
}

func (s *meta) APIKey() string {
	return s.apiKey
}

func (s *meta) APISecret() string {
	return s.apiSecret
}

func (s *meta) IsProduction() bool {
	return s.isProduction
}

func (s *meta) APIVersion() string {
	return s.apiVersion
}

func (s *meta) targetDomain() string {
	return s.domain
}

func (s *meta) setTargetDomain(n string) {
	s.domain = n
}

func (s *meta) setAPIVersion(v string) {
	s.apiVersion = v
}

// getBaseURL builds the "base" of our URL
// It builds out the host plus version. eg:
// https://api.godaddy.com/v1 for production version 1
func (s *meta) getBaseURL() (string, error) {
	if s.apiVersion == "" {
		return "", errors.New("API version not present")
	}

	url := "https://api-ote.godaddy.com/" // Development
	if s.isProduction {
		url = "https://api.godaddy.com/" // Production
	}

	return url + s.apiVersion, nil
}
*/
