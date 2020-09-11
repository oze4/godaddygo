package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request holds request data
type Request struct {
	// GoDaddy API Key, note that the prod and dev API's have unique API keys/secrets
	APIKey string
	// GoDaddy API Secret, note that the prod and dev API's have unique API keys/secrets
	APISecret string
	// HTTP REST method we validate this
	Method string
	// The URL you wish to send your request to
	URL string
	// The GoDaddy domain name you wish to target - mostly used internally
	Host string
	// The body of your request, if you need one
	Body []byte
}

// Send is sends our http request
// By default, this func does the following:
//  - Validates REST method
//  - Adds appropriate GoDaddy authorization header
//  - Sets `Content-Type` header to `application/json`
func (r *Request) Send() ([]byte, error) {
	// Verify we were given a valid REST method
	if valid := validate(r.Method, RequestMethods); valid != true {
		return nil, fmt.Errorf("Invalid request method: %s", r.Method)
	}

	// Sort out whether or not there is a Body
	var bodyFin io.ReadCloser
	if r.Body != nil {
		bodyFin = ioutil.NopCloser(strings.NewReader(string(r.Body)))
	}

	// Create new REST request
	req, err := http.NewRequest(r.Method, r.URL, bodyFin)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %s", err.Error())
	}

	// Add authorization & content-type headers to our request
	req.Header.Set("Authorization", r.makeAuthString())
	req.Header.Set("Content-Type", "application/json")

	// Create new http client to send our request
	httpclient := &http.Client{}

	// Send request, check for error
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %s", err.Error())
	}

	// Express intent to close body once we are through with it
	defer resp.Body.Close()

	// Read response body
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Verify http status
	if err := r.verifyStatusCode(resp, result); err != nil {
		return nil, err
	}

	// Return response body as bytes to simplify consumption
	return result, nil
}

// makeAuthString is used to generate the required auth string for the GoDaddy API
func (r *Request) makeAuthString() string {
	return "sso-key " + r.APIKey + ":" + r.APISecret
}

// verifyStatusCode ensure we got a good response
func (r *Request) verifyStatusCode(resp *http.Response, bodyBytes []byte) error {
	// If status code greater than 400 it should be an error
	// https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
	if resp.StatusCode >= 400 {
		var respMap map[string]string
		if err := json.Unmarshal(bodyBytes, &respMap); err != nil {
			return fmt.Errorf("Bad request\nStatus Code: %d\nError: %s", resp.StatusCode, err.Error())
		}

		// Outputs error object (map[string]string) as a single line
		// string in `key:value` notation. This is for legibility
		var status []string
		for k, v := range respMap {
			status = append(status, k+":"+v)
		}

		return errors.New(strings.Join(status, " || "))
	}
	return nil
}

func validate(s string, m map[string]string) bool {
	for t := range m {
		if s == t {
			return true
		}
	}
	return false
}

// RequestMethods holds all acceptable request methods
// https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods
var RequestMethods = map[string]string{
	"GET":     "GET",
	"HEAD":    "HEAD",
	"POST":    "POST",
	"DELETE":  "DELETE",
	"PATCH":   "PATCH",
	"OPTIONS": "OPTIONS",
	"CONNECT": "CONNECT",
	"PUT":     "PUT",
	"TRACE":   "TRACE",
}
