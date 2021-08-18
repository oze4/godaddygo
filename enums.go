package godaddygo

//
// APIEnv represents which endpoint to target (dev|prod)
//
type APIEnv string

func (e APIEnv) String() string {
	switch e {
	case APIProdEnv:
		return "prod"
	case APIDevEnv:
		return "dev"
	default:
		return ""
	}
}

// IsValid determines whether or not the given api env is valid
func (e APIEnv) IsValid() bool {
	return e.String() != ""
}

//
// APIVersion represents which endpoint version to target (v1|v2)
//
type APIVersion string

func (v APIVersion) String() string {
	switch v {
	case APIVersion1:
		return "v1"
	case APIVersion2:
		return "v2"
	default:
		return ""
	}
}

// IsValid determines whether or not the given api version is valid
func (v APIVersion) IsValid() bool {
	return v.String() != ""
}

//
// APIURL represents which URL to target
//
type APIURL string

func (u APIURL) String() string {
	switch u {
	case prodbaseURL:
		return "https://api.godaddy.com"
	case devbaseURL:
		return "https://api.ote-godaddy.com"
	default:
		return ""
	}
}

// IsValid determines whether or not the given api url is valid
func (u APIURL) IsValid() bool {
	return u.String() != ""
}

//
// RecordType represents a DNS record type
//
type RecordType string

func (r RecordType) String() string {
	switch r {
	case RecordTypeA:
		return "A"
	case RecordTypeAAAA:
		return "AAAA"
	case RecordTypeCNAME:
		return "CNAME"
	case RecordTypeMX:
		return "MX"
	case RecordTypeNS:
		return "NS"
	case RecordTypeSOA:
		return "SOA"
	case RecordTypeSRV:
		return "SRV"
	case RecordTypeTXT:
		return "TXT"
	default:
		return ""
	}
}

// IsValid determines whether or not the given dns recorsd type is valid
func (r RecordType) IsValid() bool {
	return r.String() != ""
}
