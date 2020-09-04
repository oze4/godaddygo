package api

import (
	"encoding/json"

	domainsEndpoint "github.com/oze4/godaddygo/pkg/endpoints/domains"
	"github.com/oze4/godaddygo/pkg/http"
	urls "github.com/oze4/godaddygo/pkg/url"
)

// Domain implements Domain [interface]
type Domain struct {
	Records Records
	currentRequest
}

// Contacts builds out the contacts piece of the URL
func (d *Domain) Contacts() Contacts {
	return &contacts{d.currentRequest}
}

// Privacy builds out the privacy piece of the URL
func (d *Domain) Privacy() Privacy {
	return &privacy{d.currentRequest}
}

// Agreements builds the agreements piece of the URL
func (d *Domain) Agreements(domains []string, privacyRequested, forTransfer bool) error {
	/*
		d.URL = d.URL + "/domains"
		doms := append(domains, d.Host)
		dl := strings.Join(doms, ",")
		p := strconv.FormatBool(privacyRequested)
		f := strconv.FormatBool(forTransfer)
		d.URL = "/agreements?tlds=" + dl + "&privacy=" + p + "&forTransfer=" + f
		return d.Request
	*/
	return nil
}

// IsAvailable checks if the supplied domain name is available for purchase
func (d *Domain) IsAvailable() (domainsEndpoint.Available, error) {
	req := &http.Request{
		APIKey:    d.apiKey,
		APISecret: d.apiSecret,
		URL:       urls.New(d.isProduction).Domain(d.domainName).IsAvailable(false),
		Method:    "GET",
	}

	res, err := req.Do()
	if err != nil {
		return domainsEndpoint.Available{}, err
	}

	var avail domainsEndpoint.Available
	if err := json.Unmarshal(res, &avail); err != nil {
		return domainsEndpoint.Available{}, err
	}

	return avail, nil
}

// GetDetails gets info on a domain
func (d *Domain) GetDetails() (domainsEndpoint.DomainDetails, error) {
	req := &http.Request{
		APIKey:    d.apiKey,
		APISecret: d.apiSecret,
		URL:       urls.New(d.isProduction).Domain(d.domainName).Details(),
		Method:    "GET",
	}

	resp, err := req.Do()
	if err != nil {
		return domainsEndpoint.DomainDetails{}, err
	}

	var details domainsEndpoint.DomainDetails
	if err := json.Unmarshal(resp, &details); err != nil {
		return domainsEndpoint.DomainDetails{}, err
	}

	return details, nil
}

// Delete deletes a domain
func (d *Domain) Delete() error {
	// d.URL = d.URL + "/domains/" + d.Host
	// d.Method = "DELETE"
	// return d.Request
	return nil
}

// Update updates a domain
func (d *Domain) Update(body []byte) error {
	// d.URL = d.URL + "/domains/" + d.Host
	// d.Method = "PATCH"
	// d.Body = body
	// return d.Request
	return nil
}
