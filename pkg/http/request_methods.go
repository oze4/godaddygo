package http

// RequestMethods holds all acceptable request methods
// https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods
var RequestMethods = map[string]string{
	"GET":     "GET",
	"HEAD":    "HEAD",
	"POST":    "POST",
	"DELETE":  "DELETE",
	"PATCH":   "PATCH",
	"OPTIONS": "OPTIONS",
	"CONNECT": "CONNECT",
	"PUT":     "PUT",
	"TRACE":   "TRACE",
}