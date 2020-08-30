package http

import (
	"fmt"
	"errors"

	"github.com/oze4/godaddygo/internal/validator"
	"github.com/valyala/fasthttp"
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
	APIKey    string
	APISecret string
	Method    string
	URL       string
	Host      string
	Body      []byte
}

// Do sends the http request
func (r *Request) Do() (bodyBytes []byte, err error) {
	valid := validator.Validate(r.Method, RequestMethods)
	if valid != true {
		return nil, errors.New("Invalid request method")
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(r.URL)
	req.Header.SetMethodBytes([]byte(r.Method))
	authStr := r.makeAuthString()
	fmt.Println(authStr)
	req.Header.Add("Authorization", authStr)
	req.Header.SetContentType("application/json")

	if err = fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes = resp.Body()
	println(string(bodyBytes))

	return bodyBytes, nil
}

func (r *Request) makeAuthString() string {
	return "sso " + r.APIKey + ":" + r.APISecret
}
