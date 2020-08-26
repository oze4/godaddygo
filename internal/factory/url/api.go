package url

// APIV1 represents GoDaddy's API version 1
//type APIV1 interface{
//    Domain(string) Domain
//    Domains() Domains
//}

// V1 represents GoDaddy's API version 1
type V1 struct {
	GoDaddy
}

// Domain is most likely what you're looking for. It allows you to modify domains you control
func (a V1) Domain(d string) Domain {
    return Domain{
		GoDaddy{
			URL: a.URL + "/domains/", 
			DomainName: d,
		},
	}
}

// Domains is meant for general stuff and does not require authentication (TODO: verify that)
// Like searching for available domains, or viewing public info on existing domains.
//func (a V1) Domains() Domains {
//	e := Domains{url: a.url + "/domains"}
//	return e
//}