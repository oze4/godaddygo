package godaddygo

const (
	// ProdEnv targets production API
	ProdEnv = APIProdEnv
	// DevEnv targets development API
	DevEnv = APIDevEnv
)

// NewProductionV1 targets version 1 of the production API
func NewProductionV1(key string, secret string) API {
	return NewAPI(key, secret, ProdEnv)
}

// NewV1WithClient targets version 1 of the GoDaddy API
// and allows you to use your own client
func NewV1WithClient(c Client) API {
	return WithClient(c)
}

/*
func NewV2(key string, secret string, env string) API {
	return v2.New(key, secret, env)
}

func V2WithClient(c Client) API {
	return v2.WithClient(c)
}
*/
