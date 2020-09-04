package session

// Interface holds session related data
type Interface interface {
	APIKey() string
	APISecret() string
	APIVersion() string
	IsProduction() bool
	TargetDomain() string
	SetTargetDomain(n string)
}
