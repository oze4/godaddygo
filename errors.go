package godaddygo

import (
	"fmt"
)

var (
	exception = &errs{}
)

type errs struct{}

// invalidStatusCode is when we recieve a bad status code from GoDaddy API
func (e *errs) invalidStatusCode(expectedStatus, gotStatus int, err error) error {
	return fmt.Errorf("ErrorInvalidStatusCode : expected %d, got %d\n%w", expectedStatus, gotStatus, err)
}

// invalidAPIVersion is the error you get when an incorrect Gateway version is privided within a config
func (e *errs) invalidAPIVersion(err error) error {
	return fmt.Errorf("ErrorInvalidAPIVersion : %w", err)
}

// invalidAPIEnv is the error you get when an incorrect Gateway env (production or development) is privided within a config
func (e *errs) invalidAPIEnv(err error) error {
	return fmt.Errorf("ErrorInvalidAPIEnv : invalid env (production or development) provided within config\n\t-Use `APIProdEnv` or `APIDevEnv`\n%w", err)
}

// readingBodyContent is thrown when we are unable to read body content
func (e *errs) readingBodyContent(err error) error {
	return fmt.Errorf("ErrorCannotReadBodyContent : %w", err)
}

// invalidJSONResponse is thrown when we are unable to read JSON response
func (e *errs) invalidJSONResponse(err error) error {
	return fmt.Errorf("ErrorInvalidJSONResponse : %w", err)
}

// sendingRequest is thrown when we are unable to send a request
func (e *errs) sendingRequest(err error) error {
	return fmt.Errorf("ErrorSendingRequest: %w", err)
}

// creatingNewRequest is thrown when we are unable to send a request
func (e *errs) creatingNewRequest(err error) error {
	return fmt.Errorf("ErrorCreatingNewRequest: %w", err)
}

// listingDomains is thrown when we are unable to list DNS records
func (e *errs) listingDomains(err error) error {
	return fmt.Errorf("ErrorCannotListDomains : %w", err)
}

// listingRecords is thrown when we are unable to list DNS records
func (e *errs) listingRecords(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotListRecords : %s\n%w", domainName, err)
}

// findingRecordsByType is thrown when we are unable to list DNS records
func (e *errs) findingRecordsByType(err error, domainName, recordType string) error {
	return fmt.Errorf("ErrorCannotFindRecords : byType %s of %s\n%w", recordType, domainName, err)
}

// findingRecordsByTypeAndName is thrown when we are unable to list DNS records
func (e *errs) findingRecordsByTypeAndName(err error, domainName, recordType, recordName string) error {
	return fmt.Errorf("ErrorCannotFindRecords : byType %s andName %s of %s\n%w", recordType, recordName, domainName, err)
}

// purchasingDomain is thrown when we are unable to list DNS records
func (e *errs) purchasingDomain(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotPurchaseDomain : %s\n%w", domainName, err)
}

// checkingAvailability is thrown when we are unable to check domain availability
func (e *errs) checkingAvailability(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotCheckAvailability : %s\n%w", domainName, err)
}

// gettingDomainDetails is thrown when there is an error getting domain details
func (e *errs) gettingDomainDetails(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotGetDomainDetails : %s\n%w", domainName, err)
}

// updatingRecord is thrown when an an error is encountered updating a DNS record
func (e *errs) updatingRecord(err error, domainName string, recordName string) error {
	return fmt.Errorf("ErrorUpdatingRecord : record %s of %s\n%w", recordName, domainName, err)
}
