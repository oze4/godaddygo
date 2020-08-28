package api

// Gateway allows you to select different versions of the GoDaddy API
type Gateway interface {
	V1Getter
}

// versions implements APIInterface
type gateway struct {
	*request
}

// V1 returns the V1 section of the GoDaddy API
func (a *gateway) V1() V1Interface {
	a.url = a.url + "/v1"
	return &v1{a.request}
}
