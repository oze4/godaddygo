package api

import (
	"encoding/json"
	"errors"

	"github.com/oze4/godaddygo/internal/validator"
	domainsEndpoint "github.com/oze4/godaddygo/pkg/endpoints/domains"
	"github.com/oze4/godaddygo/pkg/http"
)

// RecordsGetter makes embedding easier
type RecordsGetter interface {
	Records() Records
}

// Records allows you to interact with DNS records for a
// specific domain. Get or set DNS records
type Records interface {
	GetAll() (*[]domainsEndpoint.DNSRecord, error)
}

// records implements Records
type records struct {
	*http.Request
	recordType string
	recordName string
}

// GetAll returns all DNS records for a specific domain
func (r *records) GetAll() (*[]domainsEndpoint.DNSRecord, error) {
	r.Method = "GET"

	// In `records`, attaching should be done
	// just prior to sending request
	r.attach()

	var dnsrecords []domainsEndpoint.DNSRecord
	if err := r.getRecords(&dnsrecords, false); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// GetByType returns all DNS records for a specific domain
func (r *records) GetByType(recordType string) (*[]domainsEndpoint.DNSRecord, error) {
	r.recordType = recordType
	r.Method = "GET"

	// In `records`, attaching should be done
	// just prior to sending request
	r.attachType()

	var dnsrecords []domainsEndpoint.DNSRecord
	if err := r.getRecords(&dnsrecords, true); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// GetByTypeName allows you to get specific DNS record(s) by type and name
func (r *records) GetByTypeName(recordType, recordName string) (*[]domainsEndpoint.DNSRecord, error) {
	r.recordType = recordType
	r.recordName = recordName
	r.Method = "GET"

	// In `records`, attaching should be done
	// just prior to sending request
	r.attachTypeName()

	var dnsrecords []domainsEndpoint.DNSRecord
	if err := r.getRecords(&dnsrecords, true); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// attach appends appropriate data to URL
// string (just the `/records` piece)
func (r *records) attach() {
	r.URL = r.URL + "/records"
}

// attachType appends appropriate data to URL
// string (`/records` plus type, eg: `/records/A`)
func (r *records) attachType() {
	r.attach()
	r.URL = r.URL + "/" + r.recordType
}

// attachTypeName appends appropriate data to URL
// string (`/records` plus type plus name, eg: `/records/A/api`)
func (r *records) attachTypeName() {
	r.attachType()
	r.URL = r.URL + "/" + r.recordName
}

// getRecords is the workhorse of `records` - it is a wrapper around
// the existing request, as to not duplicate code. We expect "key"
// props to be set elsewhere, before calling this func (like r.URL
// for example). Argument `shouldValidateRecordType` calls
// `r.validateRecordType` under the hood
func (r *records) getRecords(unmarshalTo interface{}, shouldValidateRecordType bool) error {
	if shouldValidateRecordType == true {
		// Check we were given a valid record type (A, AAAA, etc....)
		if err := r.validateRecordType(); err != nil {
			return err
		}
	}

	// Since all other properties of this request are set before
	// this func is called, we can go ahead and send the request

	// Send our request, check for errors
	resp, err := r.Request.Do()
	if err != nil {
		return err
	}

	// Unmarshal the response and return data
	if err := json.Unmarshal(resp, &unmarshalTo); err != nil {
		return err
	}

	return nil
}

// validateRecordType ensures we were given an acceptable DNS record type
func (r *records) validateRecordType() error {
	if valid := validator.Validate(r.recordType, domainsEndpoint.DNSRecordTypes); valid != true {
		return errors.New("Invalid DNS Record type specified: " + r.recordType + "'")
	}
	return nil
}
