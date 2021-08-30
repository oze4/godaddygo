package godaddygo

// Contact defines the details of a contact
// Contact.Phone should be in the following format : "+1.9993335555"
type Contact struct {
	AddressMailing AddressMailing `json:"addressMailing,omitempty"`
	Email          string         `json:"email,omitempty"`
	Fax            string         `json:"fax,omitempty"`
	JobTitle       string         `json:"jobTitle,omitempty"`
	NameFirst      string         `json:"nameFirst,omitempty"`
	NameLast       string         `json:"nameLast,omitempty"`
	NameMiddle     string         `json:"nameMiddle,omitempty"`
	Organization   string         `json:"organization,omitempty"`
	Phone          string         `json:"phone,omitempty"`
}