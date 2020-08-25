package domains

// Address is a struct that holds physical address info
type Address struct {
	Address1 string `json:"address"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
}
