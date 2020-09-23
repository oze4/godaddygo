package godaddygo

// newV2 is for internal convenience
func newV2(c *Config) *v2 {
	c.version = APIVersion2
	return &v2{c}
}

// v2 implements V2
type v2 struct {
	c *Config
}
