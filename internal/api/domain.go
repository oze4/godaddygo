package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/oze4/godaddygo/pkg/endpoints/domains"
	"github.com/valyala/fasthttp"
)

// DomainGetter returns
type DomainGetter interface {
	Domain(string) DomainInterface
}

// DomainInterface represents the `domains` GoDaddy API endpoint
type DomainInterface interface {
	Contacts() Contacts
	Privacy() Privacy
	Agreements([]string, bool, bool) string
	Available() string
	Records() Records
	GetDetails() (domains.DomainDetails, error)
}

// domain implements DomainInterface
type domain struct {
	url  string
	name string
}

// Contacts builds out the contacts piece of the URL
func (d *domain) Contacts() Contacts {
	// return d.url + "/contacts"
	return Contacts{url: d.url}
}

// Privacy builds out the privacy piece of the URL
func (d *domain) Privacy() Privacy {
	// return Privacy{url: d.url + "/privacy"}
	return Privacy{url: d.url + "/privacy"}
}

// Agreements builds the agreements piece of the URL
func (d *domain) Agreements(domains []string, privacyRequested, forTransfer bool) string {
	dl := strings.Join(domains, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	return d.url + "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
}

// Available builds the available piece of the URL
func (d *domain) Available() string {
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	return d.url + "/available?domain=" + d.name + "&checkType=FAST&forTransfer=false"
}

// Records builds the DNS record piece of the URL
func (d *domain) Records() Records {
	return Records{}
}

// GetDetails gets info on a domain
func (d *domain) GetDetails() (details domains.DomainDetails, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.Add("Authorization", "sso "+"UPDATE_API_KEY"+":"+"UPDATE_API_SECRET")
	// fasthttp does not automatically request a gzipped response. We must explicitly ask for it.
	// req.Header.Set("Accept-Encoding", "gzip")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		m := fmt.Sprintf("Client get failed: %s\n", err)
		return details, errors.New(m)
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		m := fmt.Sprintf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
		return details, errors.New(m)
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

	dd := &domains.DomainDetails{}
	if err = json.Unmarshal(body, dd); err != nil {
		return details, err
	}

	return *dd, nil
}

// Delete deletes a domain
func (d *domain) Delete() {
	//TODO: Delete logic here
}

// Update updates a domain
func (d *domain) Update() {
	//TODO: Update logic here
}
