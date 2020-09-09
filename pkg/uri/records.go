package uri

// Records is the `/domain/<domain>/records` endpoint
type Records interface {
	ByType(rectype string) string
	ByTypeName(rectype, recname string) string
	String() string
}

type records struct {
	*cache
}

func (r *records) String() string {
	return r.path
}

func (r *records) ByType(rectype string) string {
	return r.path + "/" + rectype
}

func (r *records) ByTypeName(rectype, recname string) string {
	return r.ByType(rectype) + "/" + recname
}
