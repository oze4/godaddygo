package url

// APIV1 represents GoDaddy's API version 1
//type APIV1 interface{
//    Domain(string) Domain
//    Domains() Domains
//}

// APIV1 represents GoDaddy's API version 1
type APIV1 struct {
	url string
}

// Domain is most likely what you're looking for. It allows you to modify domains you control
func (a APIV1) Domain(d string) Domain {
    e := Domain{url: a.url + "/domains/" + d}
    return e
}

// Domains is meant for general stuff and does not require authentication (TODO: verify that)
// Like searching for available domains, or viewing public info on existing domains.
func (a APIV1) Domains() Domains {
	e := Domains{url: a.url + "/domains"}
	return e
}
