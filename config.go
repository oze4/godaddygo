package godaddygo

import (
    "net/http"
)

// Config holds connection options. Use `.NewConfig` to create a new config
type Config struct {
	client     *http.Client
	key        string     // key is the api key
	secret     string     // secret is the api secret
	baseURL    APIURL     // we take care of this
	env        APIEnv     // env is whether or not we are targeting prod or dev, use APIDevEnv or APIProdEnv
	version    APIVersion // version should be `v1` or `v2`, use APIVersion1 or APIVersion2
	domainName string     // dns zone to target
}