package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// InitProduction targets GoDaddy's production API (https://api.godaddy.com)
func InitProduction(r *http.Request, key, secret string) Gateway {
	return &gateway{r}
}

// InitDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func InitDevelopment() Gateway {
	panic("The OTE (development) section of this library is under construction!")
	// return &api{url: "https://api.ote-godaddy.com"}
}