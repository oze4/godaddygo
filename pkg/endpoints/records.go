package endpoints

import (
	"encoding/json"
	"errors"

	"github.com/oze4/godaddygo/internal/http"
)

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

// Records implements Records
type Records interface {
	GetAll() (*DNSRecord, error)
	GetByType(recordType string) error
	GetByTypeName(recordType, recordName string) error
}

type records struct {
	connectionBridge
}

// GetAll returns all DNS records for a specific domain
func (r *records) GetAll() (*DNSRecord, error) {
	url, err := r.getBaseURL()
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		APIKey:    r.APIKey(),
		APISecret: r.APISecret(),
		Method:    "GET",
		URL:       url + "/domains/" + r.targetDomain() + "/records",
	}

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}

	var dnsrecords DNSRecord
	if err := json.Unmarshal(resp, &dnsrecords); err != nil {
		return nil, err
	}

	return &dnsrecords, nil
}

// GetByType returns all DNS records for a specific domain
func (r *records) GetByType(recordType string) /* DNSRecord,*/ error {
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

		var dnsrecords DNSRecord
		if err := json.Unmarshal(resp, &dnsrecords); err != nil {
			return nil, err
		}

		return &dnsrecords, nil
	*/
	return nil
}

// GetByTypeName allows you to get specific DNS record(s) by type and name
func (r *records) GetByTypeName(recordType, recordName string) /* DNSRecord, */ error {
	// Check we were given a valid record type (A, AAAA, etc....)
	/*
		    if err := validateRecordType(recordType); err != nil {
					return nil, err
				}

				r.URL = r.URL + "/records/" + recordType + "/" + recordName
				r.Method = "GET"

				resp, err := r.Request.Do()
				if err != nil {
					return nil, err
				}

				var dnsrecords DNSRecord
				if err := json.Unmarshal(resp, &dnsrecords); err != nil {
					return nil, err
				}

			    return &dnsrecords, nil
	*/
	return nil
}

// SetValue allows you to set the value of an existing DNS record.
// If some.example.com resolved to 1.1.1.1 but I wanted it to be 2.2.2.2
// I would use this function to update that value
func (r *records) SetValue(recType, recName, newValue string) error {
	/*
		    // Check we were given a valid record type (A, AAAA, etc....)
			if err := validateRecordType(recType); err != nil {
				return err
			}

			newdns := DNSRecord{
				DNSRecord{
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
	*/
	return nil
}

// Add adds a new DNS record, it will NOT update any existing records
// a new record WILL be added
func (r *records) Add(rec *DNSRecord) error {
	/*
		    // Check we were given a valid record type (A, AAAA, etc....)
			if err := validateRecordType(rec.Type); err != nil {
				return err
			}

			newdns := DNSRecord{*rec}
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
	*/

	return nil
}

// AddMultiple lets you add multiple DNS records at once, it will NOT
// update any existing records a new record WILL be added
func (r *records) AddMultiple(recs *DNSRecord) error {
	/*
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
	*/

	return nil
}

// validateRecordType ensures we were given an acceptable DNS record type
func validateRecordType(recType string) error {
	if valid := validate(recType, DNSRecordTypes); valid != true {
		return errors.New("Invalid DNS Record type specified: " + recType + "'")
	}
	return nil
}
