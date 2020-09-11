package endpoints

import (
	"encoding/json"
	"errors"
	"strings"
)

// New lets you build a new record
func newRecords(s *session) Records {
	return &records{s}
}

// RecordsGetter simplifies embedding
type RecordsGetter interface {
	Records() Records
}

// Records defines `records` behavior
type Records interface {
	Add(rec *DNSRecord) error
	AddMultiple(recordsToAdd *[]DNSRecord) error
	GetAll() (*[]DNSRecord, error)
	GetByType(recordType string) (*[]DNSRecord, error)
	GetByTypeName(recordType, recordName string) (*[]DNSRecord, error)
	SetValue(recType, recName, newValue string) error
}

type records struct {
	*session
}

// GetAll returns all DNS records for a specific domain
func (r *records) GetAll() (*[]DNSRecord, error) {
	r.URL = r.URLBuilder().Domain(r.domainName).Records().String()
	r.Method = "GET"

	resp, err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	var dnsrecords []DNSRecord
	if err := json.Unmarshal(resp, &dnsrecords); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// GetByType returns all DNS records for a specific domain
func (r *records) GetByType(recordType string) (*[]DNSRecord, error) {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(recordType); err != nil {
		return nil, err
	}

	r.Method = "GET"
	r.URL = r.URLBuilder().Domain(r.domainName).Records().ByType(recordType)

	res, err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	var dnsrecords []DNSRecord
	if err := json.Unmarshal(res, &dnsrecords); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// GetByTypeName allows you to get specific DNS record(s) by type and name
func (r *records) GetByTypeName(recordType, recordName string) (*[]DNSRecord, error) {
	// Check we were given a valid record type (A, AAAA, etc....)
	if err := validateRecordType(recordType); err != nil {
		return nil, err
	}

	r.Method = "GET"
	r.URL = r.URLBuilder().Domain(r.domainName).Records().ByTypeName(recordType, recordName)

	res, err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	var dnsrecords []DNSRecord
	if err := json.Unmarshal(res, &dnsrecords); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// SetValue allows you to set the value of an existing DNS record.
// If some.example.com resolved to 1.1.1.1 but I wanted it to be 2.2.2.2
// I would use this function to update that value
func (r *records) SetValue(recType, recName, newValue string) error {
	if err := validateRecordType(recType); err != nil {
		return err
	}

	// We allow the user to just supply a new value as a string
	// This means we have to create the DNS record object which
	// the GoDaddy api expects.
	// They are looking for an array, even if it's just for a
	// single record
	rec := []DNSRecord{
		DNSRecord{
			Type: recType,
			Name: recName,
			Data: newValue,
		},
	}

	// Marshal and send
	newrec, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	r.Method = "PUT"
	r.URL = r.URLBuilder().Domain(r.domainName).Records().SetValue(recType, recName)
	r.Body = newrec

	if _, err := r.Request.Send(); err != nil {
		return err
	}

	return nil
}

// Add adds a new DNS record, it will NOT update any existing records
// a new record WILL be added
func (r *records) Add(rec *DNSRecord) error {
	if err := validateRecordType(rec.Type); err != nil {
		return err
	}

	newrec, err := json.Marshal(rec)
	if err != nil {
		return err
	}

	r.Method = "PATCH"
	r.URL = r.URLBuilder().Domain(r.domainName).Records().String()
	r.Body = newrec

	if _, err = r.Request.Send(); err != nil {
		return err
	}

	return nil
}

// AddMultiple lets you add multiple DNS records at once, it will NOT
// update any existing records a new record WILL be added
func (r *records) AddMultiple(recs *[]DNSRecord) error {
	var recordsWithError []string
	for _, rec := range *recs {
		if err := validateRecordType(rec.Type); err != nil {
			recordsWithError = append(recordsWithError, rec.Name)
		}
	}

	// If we have any records with a bad record type
	if len(recordsWithError) > 0 {
		errRecs := strings.Join(recordsWithError, ",")
		return errors.New("Invalid record type found on the following records: " + errRecs)
	}

	// Otherwise, marshal the records and send our request
	newrecs, err := json.Marshal(recs)
	if err != nil {
		return err
	}

	r.Method = "PATCH"
	r.URL = r.URLBuilder().Domain(r.domainName).Records().String()
	r.Body = newrecs

	if _, err = r.Request.Send(); err != nil {
		return err
	}

	return nil
}

// validateRecordType ensures we were given an acceptable DNS record type
func validateRecordType(recType string) error {
	if valid := validate(recType, DNSRecordTypes); valid != true {
		return errors.New("Invalid DNS Record type specified: " + recType + "'")
	}
	return nil
}

func validate(s string, m map[string]string) bool {
	for t := range m {
		if s == t {
			return true
		}
	}
	return false
}

// DNSRecord is a struct that holds data about DNS records
type DNSRecord struct {
	Data     string `json:"data,omitempty"`
	Name     string `json:"name,omitempty"`
	Port     int    `json:"port,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Service  string `json:"service,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	Type     string `json:"type,omitempty"`
	Weight   int    `json:"weight,omitempty"`
}

// DNSRecordTypes to be used as an enum
var DNSRecordTypes = map[string]string{
	"A":     "A",
	"AAAA":  "AAAA",
	"CNAME": "CNAME",
	"MX":    "MX",
	"NS":    "NS",
	"SOA":   "SOA",
	"SRV":   "SRV",
	"TXT":   "TXT",
}
