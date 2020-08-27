package http

import (
	"github.com/oze4/godaddygo/internal/validator"
)

// MethodTypes hold acceptable method types
var MethodTypes = map[string]string{
	"GET":    "GET",
	"POST":   "POST",
	"PUT":    "PUT",
	"PATCH":  "PATCH",
	"DELETE": "DELETE",
}

// ValidateMethod validates userland request methods
func ValidateMethod(m string) bool {
	return validator.Validate(m, MethodTypes)
}