package godaddygo

import (
	"fmt"
)

var (
	// ErrorWrongStatusCode is the error generated when an incorrect http status code is received
	ErrorWrongStatusCode = fmt.Errorf("ErrorWrongStatusCode")
	// ErrorWrongAPIVersion is the error you get when an incorrect API version is privided within
	// a config
	ErrorWrongAPIVersion = fmt.Errorf("ErrorWrongAPIVersion")
)
