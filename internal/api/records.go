package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// RecordsGetter makes embedding easier
type RecordsGetter interface {
	Records() Records
}

// Records is a struct builds out the `records` piece of GoDaddy's API
type Records interface {
	GetAll() *http.Request
}

type records struct {
	*http.Request
}

// func (r Records) Type() Type {}

func (r *records) GetAll() *http.Request {
	r.attach()
	r.Method = "GET"
	return r.Request
}

func (r *records) attach() {
	r.URL = r.URL + "/records"
}
