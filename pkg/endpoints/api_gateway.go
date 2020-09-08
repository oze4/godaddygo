package endpoints

import (
    "github.com/oze4/godaddygo/pkg/endpoints/versions"
)

// APIGateway targets specific GoDaddy API versions
type APIGateway interface {
    V1() versions.V1
    V2() versions.V2
}