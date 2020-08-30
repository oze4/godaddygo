package http

import (
	"errors"

	"github.com/valyala/fasthttp"
	"github.com/oze4/godaddygo/pkg/endpoints/domains"
	"github.com/oze4/godaddygo/internal/validator"
)

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
	valid := validator.Validate(r.Method, domains.DNSRecordTypes)
	if valid != true {
		return nil, errors.New("Invalid request method")
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(r.URL)
	req.Header.SetMethodBytes([]byte(r.Method))
	req.Header.Set("Authorization", r.makeAuthString())

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
