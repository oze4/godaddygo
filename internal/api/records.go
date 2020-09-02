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
	GetByType(string) (*[]domainsEndpoint.DNSRecord, error)
	GetByTypeName(string, string) (*[]domainsEndpoint.DNSRecord, error)
	SetValue(recordType string, recordName string, newValue string) error
	SetValueReturnRequest(recType, recName, newValue string) (*http.Request, error)
	Add(*domainsEndpoint.DNSRecord) error
	AddMultiple(*[]domainsEndpoint.DNSRecord) error
}

// records implements Records
type records struct {
	*http.Request
}

// GetAll returns all DNS records for a specific domain
func (r *records) GetAll() (*[]domainsEndpoint.DNSRecord, error) {
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
}

// GetByType returns all DNS records for a specific domain
func (r *records) GetByType(recordType string) (*[]domainsEndpoint.DNSRecord, error) {
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
}

// GetByTypeName allows you to get specific DNS record(s) by type and name
func (r *records) GetByTypeName(recordType, recordName string) (*[]domainsEndpoint.DNSRecord, error) {
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

// SetValueReturnRequest is for debugging purposes
func (r *records) SetValueReturnRequest(recType, recName, newValue string) (*http.Request, error) {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(recType); err != nil {
		return nil, errors.New("Invalid record type: " + recType + " " + err.Error())
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
		return nil, err
	}

	r.URL = r.URL + "/records/" + recType + "/" + recName
	r.Method = "PUT"
	r.Body = newdnsByte

	return r.Request, nil
}

// SetValue allows you to set the value of an existing DNS record.
// If some.example.com resolved to 1.1.1.1 but I wanted it to be 2.2.2.2
// I would use this function to update that value
func (r *records) SetValue(recType, recName, newValue string) error {
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
func (r *records) Add(rec *domainsEndpoint.DNSRecord) error {
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
func (r *records) AddMultiple(recs *[]domainsEndpoint.DNSRecord) error {
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
