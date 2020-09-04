package session

// Interface holds session related data
type Interface interface {
	APIKey() string
	APISecret() string
	IsProduction() bool
}
