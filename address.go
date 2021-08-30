package godaddygo

// AddressMailing defines a mailing address
type AddressMailing struct {
	Address    string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State      string `json:"state,omitempty"`
}

// FullAddress returns the full address (Address + Address2)
func (am *AddressMailing) FullAddress() string {
	return am.Address + " " + am.Address2
}
