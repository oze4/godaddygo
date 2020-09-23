package godaddygo

func newGateway(c *Config) *gateway {
	return &gateway{c}
}

type gateway struct {
	c *Config
}

func (g *gateway) V1() V1 {
	return newV1(g.c)
}

func (g *gateway) V2() V2 {
	return newV2(g.c)
}
