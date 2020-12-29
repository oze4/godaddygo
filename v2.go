package godaddygo

// newV2 is for internal convenience
func newV2(config *Config) v2 {
	config.version = APIVersion2
	return v2{config}
}

// v2 implements V2
type v2 struct {
	config *Config
}
