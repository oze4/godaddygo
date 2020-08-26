package domains

// DNSRecord is a struct that holds data about DNS records
type DNSRecord struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Port int    `json:"port"`
	Type string `json:"dnsRecordType"`
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
