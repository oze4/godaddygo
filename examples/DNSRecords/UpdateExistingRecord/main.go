package main

import (
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "// your prod key"
	prodSecret := "// your prod secret"
	targetDomain := "dom.com"
	// Connect to production API
	api := godaddygo.ConnectProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy API
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Target `records` for this domain
	records := domain.Records()
	// Update existing record
	if err := records.SetValue("A", "echo", "1.2.3.4"); err != nil {
		panic(err.Error())
	}

	fmt.Println("Success! Updated record")
}
