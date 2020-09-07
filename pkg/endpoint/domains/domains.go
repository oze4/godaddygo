package domains

import (
	"encoding/json"
	"time"

	rex "github.com/oze4/godaddygo/pkg/endpoint/domains/records"
	resthttp "github.com/oze4/godaddygo/internal/http"
)

// Domain implements Domain [interface]
type Domain interface {
	Contacts() Contacts
	Privacy() Privacy
	Agreements(domains []string, privacyRequested, forTransfer bool) error
	IsAvailable() (*Available, error)
	GetDetails() (DomainDetails, error)
	rex. Records() Records
}

type domain struct {
	meta
}

func (d *domain) Records() Records {
	return &records{d.meta}
}

// Contacts builds out the contacts piece of the URL
func (d *domain) Contacts() Contacts {
	return &contacts{d.meta}
}

// Privacy builds out the privacy piece of the URL
func (d *domain) Privacy() Privacy {
	return &privacy{d.meta}
}

// Agreements builds the agreements piece of the URL
func (d *domain) Agreements(domains []string, privacyRequested, forTransfer bool) error {
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
func (d *domain) IsAvailable() (*Available, error) {
	url, err := d.getBaseURL()
	if err != nil {
		return nil, err
	}

	dom := d.targetDomain()
	ft := "false"

	req := &resthttp.Request{
		APIKey:    d.APIKey(),
		APISecret: d.APISecret(),
		URL:       url + "/domains/available?domain=" + dom + "&checkType=FAST&forTransfer=" + ft,
		Method:    "GET",
	}

	res, err := req.Do()
	if err != nil {
		return nil, err
	}

	var avail Available
	if err := json.Unmarshal(res, &avail); err != nil {
		return nil, err
	}

	return &avail, nil
}

// GetDetails gets info on a domain
func (d *domain) GetDetails() (DomainDetails, error) {
	req := &resthttp.Request{
		APIKey:    d.APIKey(),
		APISecret: d.APISecret(),
		URL:       "", //TODO
		Method:    "GET",
	}

	resp, err := req.Do()
	if err != nil {
		return DomainDetails{}, err
	}

	var details DomainDetails
	if err := json.Unmarshal(resp, &details); err != nil {
		return DomainDetails{}, err
	}

	return details, nil
}

// Update updates a domain
// func (d *domain) Update(body []byte) error {
// d.URL = d.URL + "/domains/" + d.Host
// d.Method = "PATCH"
// d.Body = body
// return d.Request
// return nil
// }
