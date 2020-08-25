package godaddygo

// Address is a struct that holds physical address info
type Address struct {
	Address1 string `json:"address"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
}

// Consent is a struct that does something that I need to read more about
type Consent struct {
	AgreedAt      string   `json:"agreedAt"`
	AgreedBy      string   `json:"agreedBy"`
	AgreementKeys []string `json:"agreementKeys"`
}

// Contact holds contact information
type Contact struct {
	AddressMailing AddressMailing `json:"addressMailing"`
	Email          string         `json:"email"`
	Fax            string         `json:"fax"`
	JobTitle       string         `json:"jobTitle"`
	NameFirst      string         `json:"nameFirst"`
	NameLast       string         `json:"nameLast"`
	NameMiddle     string         `json:"nameMiddle"`
	Organization   string         `json:"organization"`
	Phone          string         `json:"phone"`
}

// AddressMailing is a property on Contact
type AddressMailing interface{}

// DNSRecord is a struct that holds data about DNS records
type DNSRecord struct {
	Data string        `json:"data"`
	Name string        `json:"name"`
	Port int           `json:"port"`
	Type DNSRecordType `json:"dnsRecordType"`
}

// DNSRecordType is used in DNSRecord
type DNSRecordType int

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
