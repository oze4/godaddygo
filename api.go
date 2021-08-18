package godaddygo

import (
	"github.com/oze4/godaddygo/internal/exception"
)

func newAPI(config *Config) (api, error) {
	if !config.env.IsValid() {
		return api{}, exception.InvalidAPIEnv(nil)
	}
	config.baseURL = prodbaseURL
	if config.env == APIDevEnv {
		config.baseURL = devbaseURL
	}
	return api{config}, nil
}

type api struct {
	config *Config
}

func (a api) V1() V1 {
	return newV1(a.config)
}

func (a api) V2() V2 {
	return newV2(a.config)
}
