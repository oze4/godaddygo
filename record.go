package godaddygo

// Record defines a DNS record
type Record struct {
	Data     string     `json:"data,omitempty"`
	Name     string     `json:"name,omitempty"`
	Port     int        `json:"port,omitempty"`
	Priority int        `json:"priority,omitempty"`
	Protocol string     `json:"protocol,omitempty"`
	Service  string     `json:"service,omitempty"`
	TTL      int        `json:"ttl,omitempty"`
	Type     RecordType `json:"type,omitempty"`
	Weight   int        `json:"weight,omitempty"`
}