package godaddygo

import (
    godaddy "github.com/oze4/godaddygo/pkg/client"
)

// TestingFunc is for testing
//func TestingFunc() {
//    u := url.NewBuilder().APIV1().Domains().Available("xyz.com")
//    fmt.Println(u)
//}

// NewClient creates a new GoDaddy client
func NewClient() godaddy.Client {
    opts := godaddy.NewOptions()
    return godaddy.Client{Options: opts}
}