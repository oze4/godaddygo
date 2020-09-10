package client

// Interface defines how a client should behave
// By satisfying this interface, you can use your
// own client (eg: `endpoints.Connect( <yourClient> )`
type Interface interface {
	APIKey() string
	APISecret() string
	IsProduction() bool
}
