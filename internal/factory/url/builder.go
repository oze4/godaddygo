package url

// Builder interface
//type Builder interface{
//    APIV1() APIV1
//}

// Builder does a thing
type Builder struct{
    URL string
}

// NewBuilder returns a new URLBuilder struct
func NewBuilder() Builder {
	return Builder{}
}

// APIV1 return api v1 endpoints
func (u Builder) APIV1() APIV1 {
	return APIV1{url: "https://api.godaddy.com/v1"}
}
