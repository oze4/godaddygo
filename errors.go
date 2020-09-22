package godaddygo

import (
	"fmt"
)

var (
	// ErrorWrongStatusCode is the error generated when an incorrect http status code is received
	ErrorWrongStatusCode = fmt.Errorf("ErrorWrongStatusCode")
	// ErrorWrongAPIVersion is the error you get when an incorrect Gateway version is privided within
	// a config
	ErrorWrongAPIVersion = fmt.Errorf("ErrorWrongAPIVersion")
	// ErrorWrongAPIEnv is the error you get when an incorrect Gateway env (production or development)
	// is privided within a config
	ErrorWrongAPIEnv = fmt.Errorf("ErrorWrongAPIEnv: incorrect Gateway env (production or development) privided within config")
)
