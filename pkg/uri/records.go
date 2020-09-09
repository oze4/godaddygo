package uri

// Records is the `/domain/<domain>/records` endpoint
type Records interface {
	String() string
	ByType(rectype string) string
	ByTypeName(rectype, recname string) string
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
