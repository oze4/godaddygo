package uri

// Builder is the start of generating an API URL string
func Builder(isproduction bool) Gateway {
	h := "https://api.ote-godaddy.com"
	if isproduction {
		h = "https://api.godaddy.com"
	}

	return &gateway{&cache{h}}
}
