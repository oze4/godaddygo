package godaddygo

// NewProdV1 connects you to production version 1 of the GoDaddy API
func NewProdV1(key, secret string) V1 {
    return newV1(NewConfig(key, secret, APIProdEnv))
}