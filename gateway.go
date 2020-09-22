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