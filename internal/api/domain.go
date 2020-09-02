package api

import (
	"encoding/json"
	"strconv"
	"strings"

	domainsEndpoint "github.com/oze4/godaddygo/pkg/endpoints/domains"
	"github.com/oze4/godaddygo/pkg/http"
)

// DomainGetter returns
type DomainGetter interface {
	Domain(string) Domain
}

// Domain represents the `domains` GoDaddy API endpoint
type Domain interface {
	ContactsGetter
	PrivacyGetter
	RecordsGetter
	Agreements([]string, bool, bool) http.Request
	IsAvailable() (*domainsEndpoint.Available, error)
	GetDetails() (*domainsEndpoint.DomainDetails, error)
}

// domain implements Domain [interface]
type domain struct {
	http.Request
}

// Contacts builds out the contacts piece of the URL
func (d *domain) Contacts() Contacts {
	d.URL = d.URL + "/domains/" + d.Host
	return &contacts{d.Request}
}

// Privacy builds out the privacy piece of the URL
func (d *domain) Privacy() Privacy {
	d.URL = d.URL + "/domains/" + d.Host
	return &privacy{d.Request}
}

// Agreements builds the agreements piece of the URL
func (d *domain) Agreements(domains []string, privacyRequested, forTransfer bool) http.Request {
	d.URL = d.URL + "/domains"
	doms := append(domains, d.Host)
	dl := strings.Join(doms, ",")
	p := strconv.FormatBool(privacyRequested)
	f := strconv.FormatBool(forTransfer)
	d.URL = "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
	return d.Request
}

// IsAvailable checks if the supplied domain name is available for purchase
func (d *domain) IsAvailable() (*domainsEndpoint.Available, error) {
	//TODO: parameterize checkType and forTransfer in the URL (like func Agreements)
	d.URL = d.URL + "/domains/available?domain=" + d.Host + "&checkType=FAST&forTransfer=false"
	d.Method = "GET"

	res, err := d.Request.Do()
	if err != nil {
		return nil, err
	}

	var avail domainsEndpoint.Available
	if err := json.Unmarshal(res, &avail); err != nil {
		return nil, err
	}

	return &avail, nil
}

// GetDetails gets info on a domain
func (d *domain) GetDetails() (*domainsEndpoint.DomainDetails, error) {
	d.URL = d.URL + "/domains/" + d.Host
	d.Method = "GET"

	res, err := d.Request.Do()
	if err != nil {
		return nil, err
	}

	var details domainsEndpoint.DomainDetails
	if err := json.Unmarshal(res, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// Delete deletes a domain
func (d *domain) Delete() http.Request {
	d.URL = d.URL + "/domains/" + d.Host
	d.Method = "DELETE"
	return d.Request
}

// Update updates a domain
func (d *domain) Update(body []byte) http.Request {
	d.URL = d.URL + "/domains/" + d.Host
	d.Method = "PATCH"
	d.Body = body
	return d.Request
}

// Records builds the DNS record piece of the URL
func (d *domain) Records() Records {
	d.URL = d.URL + "/domains/" + d.Host
	return &records{d.Request}
}
