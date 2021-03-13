package exception

import (
	"fmt"
)

// Enum knows how to return a string representation of itself
type Enum interface {
	String() string
}

// InvalidStatusCode is when we recieve a bad status code from GoDaddy API
func InvalidStatusCode(expectedStatus, gotStatus int, err string) error {
	return fmt.Errorf("ErrorInvalidStatusCode : expected %d, got %d\n%s", expectedStatus, gotStatus, err)
}

// InvalidAPIVersion is the error you get when an incorrect Gateway version is privided within a config
func InvalidAPIVersion(err error) error {
	return fmt.Errorf("ErrorInvalidAPIVersion : %s", err.Error())
}

// InvalidAPIEnv is the error you get when an incorrect Gateway env (production or development) Is privided within a config
func InvalidAPIEnv(err error) error {
	return fmt.Errorf("ErrorInvalidAPIEnv : invalid env (production or development) provided within config\n\t-Use `APIProdEnv` or `APIDevEnv`\n%s", err.Error())
}

// ReadingBodyContent is thrown when we are unable to read body content
func ReadingBodyContent(err error) error {
	return fmt.Errorf("ErrorCannotReadBodyContent : %s", err.Error())
}

// InvalidJSONResponse is thrown when we are unable to read JSON response
func InvalidJSONResponse(err error) error {
	return fmt.Errorf("ErrorInvalidJSONResponse : %s", err.Error())
}

// SendingRequest is thrown when we are unable to send a request
func SendingRequest(err error) error {
	return fmt.Errorf("ErrorSendingRequest: %s", err.Error())
}

// CreatingNewRequest is thrown when we are unable to send a request
func CreatingNewRequest(err error) error {
	return fmt.Errorf("ErrorCreatingNewRequest: %s", err.Error())
}

// ListingDomains is thrown when we are unable to list DNS records
func ListingDomains(err error) error {
	return fmt.Errorf("ErrorCannotListDomains : %s", err.Error())
}

// ListingRecords is thrown when we are unable to list DNS records
func ListingRecords(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotListRecords : %s\n%s", domainName, err.Error())
}

// FindingRecordsByType is thrown when we are unable to list DNS records
func FindingRecordsByType(err error, domainName, recordType string) error {
	return fmt.Errorf("ErrorCannotFindRecords : byType %s of %s\n%s", recordType, domainName, err.Error())
}

// FindingRecordsByTypeAndName is thrown when we are unable to list DNS records
func FindingRecordsByTypeAndName(err error, domainName, recordType, recordName string) error {
	return fmt.Errorf("ErrorCannotFindRecords : byType %s andName %s of %s\n%s", recordType, recordName, domainName, err.Error())
}

// PurchasingDomain is thrown when we are unable to list DNS records
func PurchasingDomain(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotPurchaseDomain : %s\n%s", domainName, err.Error())
}

// CheckingAvailability is thrown when we are unable to check domain availability
func CheckingAvailability(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotCheckAvailability : %s\n%s", domainName, err.Error())
}

// GettingDomainDetails is thrown when there is an error getting domain details
func GettingDomainDetails(err error, domainName string) error {
	return fmt.Errorf("ErrorCannotGetDomainDetails : %s\n%s", domainName, err.Error())
}

// AddingRecords is thrown when an an error is encountered adding DNS record
func AddingRecords(err error, domainName string, recordName string) error {
	return fmt.Errorf("ErrorAddingRecords : record %s of %s\n%s", recordName, domainName, err.Error())
}

// UpdatingRecord is thrown when an an error is encountered updating a DNS record
func UpdatingRecord(err error, domainName string, recordName string) error {
	return fmt.Errorf("ErrorUpdatingRecord : record %s of %s\n%s", recordName, domainName, err.Error())
}

// DeletingRecord is thrown when an error is encountered deleting a DNS record
func DeletingRecord(err error, domainName string, recordName string, recordType string) error {
	return fmt.Errorf("ErrorDeletingRecord : type %s record %s of %s\n%s", recordType, recordName, domainName, err.Error())
}

// InvalidValue is thrown when an invalid value is being used for an "enum"
func InvalidValue(message string) error {
	return fmt.Errorf("invalid value : %s", message)
}
