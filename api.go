package godaddygo

import (
	"io"
	"io/ioutil"
)

func newAPI(c *Config) api {
	return api{c}
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

func bodyToBytes(body io.Reader) ([]byte, error) {
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
