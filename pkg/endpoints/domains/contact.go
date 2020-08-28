package domains

// Contact holds contact information
type Contact struct {
	AddressMailing AddressMailing `json:"addressMailing"`
	Email          string         `json:"email"`
	Fax            string         `json:"fax"`
	JobTitle       string         `json:"jobTitle"`
	NameFirst      string         `json:"nameFirst"`
	NameLast       string         `json:"nameLast"`
	NameMiddle     string         `json:"nameMiddle"`
	Organization   string         `json:"organization"`
	Phone          string         `json:"phone"`
}
