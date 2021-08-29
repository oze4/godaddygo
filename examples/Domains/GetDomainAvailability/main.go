package main

import (
	"context"
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "-"
	prodSecret := "-"
	// Connect to production Gateway
	api, err := godaddygo.NewProduction(prodKey, prodSecret)
	if err != nil {
		panic(err.Error())
	}
	// Target version 1 of the production GoDaddy Gateway
	prodv1 := api.V1()
	// Check for domain availability
	// *You DO NOT need to own this domain*
	availability, err := prodv1.CheckAvailability(context.Background(), "domtocheck.com", true)// -> *GoDaddyDomainAvailability
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(availability.Available)
}
