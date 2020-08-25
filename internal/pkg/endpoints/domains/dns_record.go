package domains

// DNSRecord is a struct that holds data about DNS records
type DNSRecord struct {
	Data string        `json:"data"`
	Name string        `json:"name"`
	Port int           `json:"port"`
	Type DNSRecordType `json:"dnsRecordType"`
}

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

func lazy() string {
	recerd := A
	return recerd.String()
}
