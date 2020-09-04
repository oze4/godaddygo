package api

import (
	"encoding/json"
	"errors"

	"github.com/oze4/godaddygo/internal/validator"
	domainsEndpoint "github.com/oze4/godaddygo/pkg/endpoints/domains"
	"github.com/oze4/godaddygo/pkg/http"
	urlFactory "github.com/oze4/godaddygo/pkg/url"
)

// Records implements Records
type Records struct {
	currentRequest
}

// GetAll returns all DNS records for a specific domain
func (r *Records) GetAll() ([]domainsEndpoint.DNSRecord, error) {
	req := &http.Request{
		Method: "GET",
		URL:    urlFactory.New(r.isProduction).Domain(r.domainName).Records.GetAll(),
	}
	/*
		r.URL = r.URL + "/records"
		r.Method = "GET"

		resp, err := r.Request.Do()
		if err != nil {
			return nil, err
		}

		var dnsrecords []domainsEndpoint.DNSRecord
		if err := json.Unmarshal(resp, &dnsrecords); err != nil {
			return nil, err
		}

		return &dnsrecords, nil
	*/
}

// GetByType returns all DNS records for a specific domain
func (r *Records) GetByType(recordType string) (*[]domainsEndpoint.DNSRecord, error) {
	/*
		// Check we were given a valid record type (A, AAAA, etc....)
		if err := validateRecordType(recordType); err != nil {
			return nil, err
		}

		r.URL = r.URL + "/records/" + recordType
		r.Method = "GET"

		resp, err := r.Request.Do()
		if err != nil {
			return nil, err
		}

		var dnsrecords []domainsEndpoint.DNSRecord
		if err := json.Unmarshal(resp, &dnsrecords); err != nil {
			return nil, err
		}

		return &dnsrecords, nil
	*/
}

// GetByTypeName allows you to get specific DNS record(s) by type and name
func (r *Records) GetByTypeName(recordType, recordName string) (*[]domainsEndpoint.DNSRecord, error) {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(recordType); err != nil {
		return nil, err
	}

	r.URL = r.URL + "/records/" + recordType + "/" + recordName
	r.Method = "GET"

	resp, err := r.Request.Do()
	if err != nil {
		return nil, err
	}

	var dnsrecords []domainsEndpoint.DNSRecord
	if err := json.Unmarshal(resp, &dnsrecords); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// SetValueReturnRequest is for debugging purposes and will be removed shortly
func (r *Records) SetValueReturnRequest(recType, recName, newValue string) (http.Request, error) {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(recType); err != nil {
		return http.Request{}, errors.New("Invalid record type: " + recType + " " + err.Error())
	}

	newdns := []domainsEndpoint.DNSRecord{
		domainsEndpoint.DNSRecord{
			Type: recType,
			Name: recName,
			Data: newValue,
		},
	}

	newdnsByte, err := json.Marshal(newdns)
	if err != nil {
		return http.Request{}, err
	}

	r.URL = r.URL + "/records/" + recType + "/" + recName
	r.Method = "PUT"
	r.Body = newdnsByte

	return r.Request, nil
}

// SetValue allows you to set the value of an existing DNS record.
// If some.example.com resolved to 1.1.1.1 but I wanted it to be 2.2.2.2
// I would use this function to update that value
func (r *Records) SetValue(recType, recName, newValue string) error {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(recType); err != nil {
		return err
	}

	newdns := []domainsEndpoint.DNSRecord{
		domainsEndpoint.DNSRecord{
			Type: recType,
			Name: recName,
			Data: newValue,
		},
	}

	newdnsByte, err := json.Marshal(newdns)
	if err != nil {
		return err
	}

	r.URL = r.URL + "/records/" + recType + "/" + recName
	r.Method = "PUT"
	r.Body = newdnsByte

	_, err = r.Request.Do()
	if err != nil {
		return err
	}

	return nil
}

// Add adds a new DNS record, it will NOT update any existing records
// a new record WILL be added
func (r *Records) Add(rec *domainsEndpoint.DNSRecord) error {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(rec.Type); err != nil {
		return err
	}

	newdns := []domainsEndpoint.DNSRecord{*rec}
	bod, err := json.Marshal(newdns)
	if err != nil {
		return err
	}

	r.URL = r.URL + "/records"
	r.Method = "PATCH"
	r.Body = bod

	if _, err = r.Request.Do(); err != nil {
		return err
	}

	return nil
}

// AddMultiple lets you add multiple DNS records at once, it will NOT
// update any existing records a new record WILL be added
func (r *Records) AddMultiple(recs *[]domainsEndpoint.DNSRecord) error {
	iserr := false
	for _, rec := range *recs {
		// Check we were given a valid record type (A, AAAA, etc....)
		if err := validateRecordType(rec.Type); err != nil {
			iserr = true
		}
	}

	if iserr {
		return errors.New("Invalid record type found")
	}

	bod, err := json.Marshal(recs)
	if err != nil {
		return err
	}

	r.URL = r.URL + "/records"
	r.Method = "PATCH"
	r.Body = bod

	if _, err = r.Request.Do(); err != nil {
		return err
	}

	return nil
}

// validateRecordType ensures we were given an acceptable DNS record type
func validateRecordType(recType string) error {
	if valid := validator.Validate(recType, domainsEndpoint.DNSRecordTypes); valid != true {
		return errors.New("Invalid DNS Record type specified: " + recType + "'")
	}
	return nil
}
