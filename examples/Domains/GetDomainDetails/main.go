package main

import (
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "-"
	prodSecret := "-"
	targetDomain := "dom.com" // You DO have to own this domain
	// Connect to production API
	api := godaddygo.ConnectProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy API
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Get details on our domain
	deets, err := domain.GetDetails() // -> *GoDaddyDomainDetails
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(deets)
}
