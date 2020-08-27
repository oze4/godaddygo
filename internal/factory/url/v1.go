package url

// V1 represents GoDaddy's API version 1
type V1 struct {
	url string
}

// Domain is most likely what you're looking for. It allows you to modify domains you control
func (v V1) Domain(d string) Domain {
    return Domain{name: d, url: v.url + "/domains/" + d}
}