package main

import (
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "-"
	prodSecret := "-"
	// Connect to production API
	api := godaddygo.ConnectProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy API
	prodv1 := api.V1()
	// Check for domain availability
	// *You DO NOT need to own this domain*
	availability, err := prodv1.GetDomainAvailability("google.com") // -> *DomainAvailability
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(availability.Available)
}
