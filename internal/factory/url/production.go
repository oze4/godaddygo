package url

// Production specifies the production base URL for the GoDaddy API (https://api.godaddy.com)
type Production struct {
	url string
}

// V1 returns the V1 section of the GoDaddy API
func (p Production) V1() V1 {
	return V1 {url: p.url + "/v1"}
}