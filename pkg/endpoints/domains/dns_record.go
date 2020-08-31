package domains

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

/*
 {
    "data": "string",
    "name": "string",
    "port": 0,
    "priority": 0,
    "protocol": "string",
    "service": "string",
    "ttl": 0,
    "type": "A",
    "weight": 0
  }
*/

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

/* TODO: move this block elsewhere - it is still useful code to learn from
// DNSRecordType is used in DNSRecord
type DNSRecordType int

// Constants to be used as an enum
const (
	A DNSRecordType = 1 + iota
	AAAA
	CNAME
	MX
	NS
	SOA
	SRV
	TXT
)

var dnsRecordTypes = [...]string{
	"A",
	"AAAA",
	"CNAME",
	"MX",
	"NS",
	"SOA",
	"SRV",
	"TXT",
}

func (drt DNSRecordType) String() string {
	return dnsRecordTypes[drt-1]
}
*/
