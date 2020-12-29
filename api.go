package godaddygo

import (
	"github.com/oze4/godaddygo/internal/exception"
)

func newAPI(config *Config) (api, error) {
	switch config.env {
	case APIProdEnv:
		config.baseURL = prodbaseURL
	case APIDevEnv:
		config.baseURL = devbaseURL
	default:
		return api{}, exception.InvalidAPIEnv(nil)
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
