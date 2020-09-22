package godaddygo

import (
    "fmt"
)

var (
	// ErrorWrongStatusCode is the error generated when an incorrect http status code is received
	ErrorWrongStatusCode = fmt.Errorf("ErrorWrongStatusCode")
)