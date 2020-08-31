package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// InitProduction targets GoDaddy's production API (https://api.godaddy.com)
func InitProduction(r *http.Request) Gateway {
	r.URL = "https://api.godaddy.com"
	return &gateway{r}
}

// InitDevelopment targets GoDaddy's development API (https://api.ote-godaddy.com)
func InitDevelopment(r *http.Request) Gateway {
	r.URL = "https://api.ote-godaddy.com"
	return &gateway{r}
}
