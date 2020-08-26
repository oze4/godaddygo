package url

// APIV1 represents GoDaddy's API version 1
type APIV1 interface{
    Domains() DomainsEndpoint
}

type apiV1 struct {
	url string
}

// Domains is meant for general stuff and does not require authentication (TODO: verify that)
// Like searching for available domains, or viewing public info on existing domains.
func (a apiV1) Domains() DomainsEndpoint {
	e := domainsEndpoint{url: a.url + "/domains"}
	return e
}
