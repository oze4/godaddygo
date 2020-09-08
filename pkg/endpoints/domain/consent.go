package domain

// Consent is a struct that does something that I need to read more about
type Consent struct {
	AgreedAt      string   `json:"agreedAt"`
	AgreedBy      string   `json:"agreedBy"`
	AgreementKeys []string `json:"agreementKeys"`
}
