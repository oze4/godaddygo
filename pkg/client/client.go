package client

import (
	"github.com/oze4/godaddygo/internal/factory/url"
)

// Client is what allows you to interact with the GoDaddy API
type Client struct {
	Options    *Options
	URLBuilder url.Builder
}
