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
	targetDomain := "dom.com" // You DO have to own this domain
	// Connect to production Gateway
	api, err := godaddygo.NewProduction(prodKey, prodSecret)
	if err != nil {
		panic(err.Error())
	}
	// Target version 1 of the production GoDaddy Gateway
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Get details on our domain
	deets, err := domain.GetDetails(context.Background()) // -> *GoDaddyDomainDetails
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(deets)
}
