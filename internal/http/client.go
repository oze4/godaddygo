package http

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

// Client holds client info
type Client struct {
	APIKey    string
	APISecret string
}

// Get performs get requests
func (c Client) Get(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.Add("Authorization", "sso "+c.APIKey+":"+c.APISecret)
	// fasthttp does not automatically request a gzipped response. We must explicitly ask for it.
	// req.Header.Set("Accept-Encoding", "gzip")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		m := fmt.Sprintf("Client get failed: %s\n", err)
		return nil, errors.New(m)
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		m := fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return nil, errors.New(m)
	}

	// Do we need to decompress the response?
	contentEncoding := resp.Header.Peek("Content-Encoding")
	var body []byte

	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		fmt.Println("Unzipping...")
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}

	return body, nil
}

// Post performs post requests
func (c Client) Post(body []byte) {

}
