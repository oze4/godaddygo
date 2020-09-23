package godaddygo

import (
	"fmt"
)

var (
	exceptions = &errs{}
)

type errs struct{}

// wrongStatusCode is when we recieve a bad status code from GoDaddy API
func (e *errs) wrongStatusCode(expectedStatus, gotStatus int) error {
	return fmt.Errorf("ErrorWrongStatusCode: expectedStatus %d, got %d", expectedStatus, gotStatus)
}

// wrongAPIVersion is the error you get when an incorrect Gateway version is privided within a config
func (e *errs) wrongAPIVersion() error {
	return fmt.Errorf("ErrorWrongAPIVersion")
}

// wrongAPIEnv is the error you get when an incorrect Gateway env (production or development) is privided within a config
func (e *errs) wrongAPIEnv() error {
	return fmt.Errorf("ErrorWrongAPIEnv: incorrect Gateway env (production or development) privided within config")
}

// cannotReadBodyContent is thrown when we are unable to read body content
func (e *errs) cannotReadBodyContent(err error) error {
	return fmt.Errorf("cannot read body content : %w", err)
}

// invalidJSONResponse is thrown when we are unable to read JSON response
func (e *errs) invalidJSONResponse(err error) error {
	return fmt.Errorf("invalid json response : %w", err)
}

// sendingRequest is thrown when we are unable to send a request
func (e *errs) sendingRequest(err error) error {
	return fmt.Errorf("Error sending request: %w", err)
}

// creatingRequest is thrown when we are unable to send a request
func (e *errs) creatingRequest(err error) error {
	return fmt.Errorf("Error creating new request: %w", err)
}

// cannotListRecords is thrown when we are unable to list DNS records
func (e *errs) cannotListRecords(domainName string, err error) error {
	return fmt.Errorf("Cannot list records of %s : %w", domainName, err)
}

// unableToPurchaseDomain is thrown when we are unable to list DNS records
func (e *errs) unableToPurchaseDomain(domainName string, err error) error {
	return fmt.Errorf("Cannot purchase domain %s : %w", domainName, err)
}

// unableToCheckAvailability is thrown when we are unable to check domain availability
func (e *errs) unableToCheckAvailability(domainName string, err error) error {
	return fmt.Errorf("Cannot get availability of domain %s : %w", domainName, err)
}