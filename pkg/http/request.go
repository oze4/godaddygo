package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	//"net/url"
	"strings"

	"github.com/oze4/godaddygo/internal/validator"
)

// RequestMethods holds all acceptable request methods
var RequestMethods = map[string]string{
	"GET":    "GET",
	"POST":   "POST",
	"DELETE": "DELETE",
	"PATCH":  "PATCH",
	"PUT":    "PUT",
}

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

// Do sends the http request
func (r *Request) Do() ([]byte, error) {
	// Verify we were given a valid REST method
	valid := validator.Validate(r.Method, RequestMethods)
	if valid != true {
		return nil, errors.New("Invalid request method")
	}

	// Sort out whether or not there is a Body
	var bodyFin io.ReadCloser
	if r.Body != nil {
		bodyFin = ioutil.NopCloser(strings.NewReader(string(r.Body)))
	}

	// Create new REST request
	req, err := http.NewRequest(r.Method, r.URL, bodyFin)
	if err != nil {
		return nil, err
	}

	// Add authorization to our request
	req.Header.Set("Authorization", r.makeAuthString())

	// Create new http client to send our request
	httpclient := &http.Client{}

	// Send request, check for error
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	// defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	// Read response body
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Close the response body
	resp.Body.Close()

	// Return response body as bytes
	return result, nil
}

// makeAuthString is used to generate the required auth string for the GoDaddy API
func (r *Request) makeAuthString() string {
	return "sso-key " + r.APIKey + ":" + r.APISecret
}
