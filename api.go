package godaddygo

func newAPI(c *Config) (api, error) {
	switch c.env {
	case APIProdEnv:
		c.baseURL = prodbaseURL
	case APIDevEnv:
		c.baseURL = devbaseURL
	default:
		return api{}, exception.invalidAPIEnv(nil)
	}

	return api{c}, nil
}

type api struct {
	c *Config
}

func (a api) V1() V1 {
	return newV1(a.c)
}

func (a api) V2() V2 {
	return newV2(a.c)
}
