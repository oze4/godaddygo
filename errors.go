package godaddygo

import (
	"fmt"
)

var (
	exceptions = &errs{}
)

type errs struct{}

// is when we recieve a bad status code from GoDaddy API
func (e *errs) errorWrongStatusCode(expectedStatus, gotStatus int) error {
	return fmt.Errorf("ErrorWrongStatusCode: expectedStatus %d, got %d", expectedStatus, gotStatus)
}

// ErrorWrongAPIVersion is the error you get when an incorrect Gateway version is privided within a config
func (e *errs) errorWrongAPIVersion() error {
	return fmt.Errorf("ErrorWrongAPIVersion")
}

// ErrorWrongAPIEnv is the error you get when an incorrect Gateway env (production or development) is privided within a config
func (e *errs) errorWrongAPIEnv() error {
	return fmt.Errorf("ErrorWrongAPIEnv: incorrect Gateway env (production or development) privided within config")
}

// ErrorCannotReadBodyContent is thrown when we are unable to read body content
func (e *errs) errorCannotReadBodyContent(err error) error {
	return fmt.Errorf("cannot read body content : %w", err)
}

// ErrorInvalidJSONResponse is thrown when we are unable to read JSON response
func (e *errs) errorInvalidJSONResponse(err error) error {
	return fmt.Errorf("invalid json response : %w", err)
}
