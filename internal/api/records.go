package api

import (
	"encoding/json"
	"github.com/oze4/godaddygo/pkg/http"
	domainsEndpoint "github.com/oze4/godaddygo/pkg/endpoints/domains"
)

// RecordsGetter makes embedding easier
type RecordsGetter interface {
	Records() Records
}

// Records is a struct builds out the `records` piece of GoDaddy's API
type Records interface {
	GetAll() (*[]domainsEndpoint.DNSRecord, error)
}

type records struct {
	*http.Request
}

// GetAll returns all DNS records for a particular domain
func (r *records) GetAll() (*[]domainsEndpoint.DNSRecord, error) {
	r.attach()
	r.Method = "GET"
	
	res, err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	var allrecords []domainsEndpoint.DNSRecord
	if err := json.Unmarshal(res, &allrecords); err != nil {
		return nil, err
	}

	return &allrecords, nil
}

func (r *records) attach() {
	r.URL = r.URL + "/records"
}
