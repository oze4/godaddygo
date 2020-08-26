package godaddygo

import (
    "fmt"
    
    "github.com/oze4/godaddygo/internal/factory/url"
)

// TestingFunc is for testing
func TestingFunc() {
    u := url.NewBuilder().APIV1().Domains().Available("xyz.com")
    fmt.Println(u)
}