package godaddygo

// DomainAvailability holds data about domain availability
type DomainAvailability struct {
	Available  bool   `json:"available,omitempty"`
	Currency   string `json:"currency,omitempty"`
	Definitive bool   `json:"definitive,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Period     int    `json:"period,omitempty"`
	Price      int    `json:"price,omitempty"`
}