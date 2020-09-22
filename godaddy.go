package godaddygo

// NewProductionV1 returns a new production v1 API
func NewProductionV1(key, secret string) *API {
	conf := NewConfig(key, secret, APIProdEnv, APIVersion1)
	return Connect(conf)
}
