package godaddygo

// NewProductionV1 returns a new production v1 API
func NewProductionV1(key, secret string) *API {
	return NewAPI(
		NewConfig(key, secret, APIProdEnv, APIVersion1),
	)
}
