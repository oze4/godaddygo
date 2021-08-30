package godaddygo

// PurchaseRequest is needed when purchasing a domain
type PurchaseRequest struct {
	Consent           Consent  `json:"consent,omitempty"`
	ContactAdmin      Contact  `json:"contactAdmin,omitempty"`
	ContactBilling    Contact  `json:"contactBilling,omitempty"`
	ContactRegistrant Contact  `json:"contactRegistrant,omitempty"`
	ContactTech       Contact  `json:"contactTech,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	NameServers       []string `json:"nameServers,omitempty"`
	Period            int      `json:"period,omitempty"`    // Min 1, Max 10, Default 1
	Privacy           bool     `json:"privacy,omitempty"`   // Default false
	RenewAuto         bool     `json:"renewAuto,omitempty"` // Default true
}

// IsValidPeriod determines if given PurchaseRequest contains a valid Period
func (p PurchaseRequest) IsValidPeriod() bool {
	return p.Period >= 1 && p.Period <= 10
}

// PurchaseReceipt is what is returned after purchasing a domain
type PurchaseReceipt struct {
	Currency  string `json:"currency,omitempty"`
	ItemCount int    `json:"itemCount,omitempty"`
	OrderID   int    `json:"orderId,omitempty"`
	Total     int    `json:"total,omitempty"`
}
