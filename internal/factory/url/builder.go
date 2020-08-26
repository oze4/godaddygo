package url

// Builder interface
type Builder interface{
    APIV1() APIV1
}

type builder struct{
    URL string
}

// NewBuilder returns a new URLBuilder struct
func NewBuilder() Builder {
	return builder{}
}

// APIV1 return api v1 endpoints
func (u builder) APIV1() APIV1 {
	return apiV1{url: "https://api.godaddy.com/v1"}
}
