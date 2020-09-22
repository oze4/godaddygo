package godaddygo

import (
    "net/http"
)

// NewConfig creates a config using `http.DefaultClient` as our client
//  - key is the api key
//  - secret is the api secret
//  - env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
//  - version should be `v1` or `v2`, use APIVersion1 or APIVersion2
func NewConfig(key, secret, env string) *Config {
	return &Config{
		Client: http.DefaultClient,
		key:    key,
		secret: secret,
		env:    env,
	}
}

// NewConfigWithClient creates a config using a custom http client
//  - key is the api key
//  - secret is the api secret
//  - env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
//  - version should be `v1` or `v2`, use APIVersion1 or APIVersion2
//  - c is your http client
func NewConfigWithClient(key, secret, env string, c *http.Client) *Config {
	return &Config{
		Client:  c,
		key:     key,
		secret:  secret,
		env:     env,
	}
}

// NewGateway creates a new GoDaddy API gateway based upon a config
func NewGateway(c *Config) Gateway {
    return newGateway(c)
}

// NewProdV1 connects you to production version 1 of the GoDaddy API
func NewProdV1(key, secret string) V1 {
    return newV1(NewConfig(key, secret, APIProdEnv))
}