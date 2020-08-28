package domains

// AddressMailing holds contact address info
type AddressMailing struct {
	Address    string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
}
