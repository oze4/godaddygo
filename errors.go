package godaddygo

import (
	"fmt"
)

var (
	exception = errs{}
)

type errs struct{}

// invalidStatusCode is when we recieve a bad status code from GoDaddy API
func (e errs) invalidStatusCode(expectedStatus, gotStatus int, err error) error {
	return fmt.Errorf("ErrorInvalidStatusCode : expected %d, got %d\n%s", expectedStatus, gotStatus, err.Error())
}

// invalidAPIVersion is the error you get when an incorrect Gateway version is privided within a config
func (e errs) invalidAPIVersion(err error) error {
	return fmt.Errorf("ErrorInvalidAPIVersion : %s", err.Error())
}

// invalidAPIEnv is the error you get when an incorrect Gateway env (production or development) is privided within a config
func (e errs) invalidAPIEnv(err error) error {
	return fmt.Errorf("ErrorInvalidAPIEnv : invalid env (production or development) provided within config\n\t-Use `APIProdEnv` or `APIDevEnv`\n%s", err.Error())
}

// readingBodyContent is thrown when we are unable to read body content
func (e errs) readingBodyContent(err error) error {
	return fmt.Errorf("ErrorCannotReadBodyContent : %s", err.Error())
}

// invalidJSONResponse is thrown when we are unable to read JSON response
func (e errs) invalidJSONResponse(err error) error {
	return fmt.Errorf("ErrorInvalidJSONResponse : %s", err.Error())
}

// sendingRequest is thrown when we are unable to send a request
func (e errs) sendingRequest(err error) error {
	return fmt.Errorf("ErrorSendingRequest: %s", err.Error())
}

// creatingNewRequest is thrown when we are unable to send a request
func (e errs) creatingNewRequest(err error) error {
	return fmt.Errorf("ErrorCreatingNewRequest: %s", err.Error())
}

// listingDomains is thrown when we are unable to list DNS records
func (e errs) listingDomains(err error) error {
	return fmt.Errorf("ErrorCannotListDomains : %s", err.Error())
}

// listingRecords is thrown when we are unable to list DNS records
func (e errs) listingRecords(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotListRecords : %s\n%s", domainName, err.Error())
}

// findingRecordsByType is thrown when we are unable to list DNS records
func (e errs) findingRecordsByType(err error, domainName, recordType string) error {
	return fmt.Errorf("ErrorCannotFindRecords : byType %s of %s\n%s", recordType, domainName, err.Error())
}

// findingRecordsByTypeAndName is thrown when we are unable to list DNS records
func (e errs) findingRecordsByTypeAndName(err error, domainName, recordType, recordName string) error {
	return fmt.Errorf("ErrorCannotFindRecords : byType %s andName %s of %s\n%s", recordType, recordName, domainName, err.Error())
}

// purchasingDomain is thrown when we are unable to list DNS records
func (e errs) purchasingDomain(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotPurchaseDomain : %s\n%s", domainName, err.Error())
}

// checkingAvailability is thrown when we are unable to check domain availability
func (e errs) checkingAvailability(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotCheckAvailability : %s\n%s", domainName, err.Error())
}

// gettingDomainDetails is thrown when there is an error getting domain details
func (e errs) gettingDomainDetails(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotGetDomainDetails : %s\n%s", domainName, err.Error())
}

// updatingRecord is thrown when an an error is encountered updating a DNS record
func (e errs) updatingRecord(err error, domainName string, recordName string) error {
	return fmt.Errorf("ErrorUpdatingRecord : record %s of %s\n%s", recordName, domainName, err.Error())
}
