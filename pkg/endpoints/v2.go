package endpoints

// newV2 creates a new v2
func newV2(s *session) V2 {
	s.apiVersion = "v1"
	return &v2{s}
}

// V2 targets version 1 of the GoDaddy API
type V2 interface {
	//TODO
}

type v2 struct {
	*session
}
