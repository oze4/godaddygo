package main

import (
	"fmt"
	godaddy "github.com/oze4/godaddygo"
)

func main() {
	MyAPIKey := "-"    // Make sure to supply appropriate key [prod or dev]
	MyAPISecret := "-" // ...same here

	// *** YOU MUST OWN THIS DOMAIN TO GET DETAILS ***
	DomainToTarget := "somedomainyouown.com"

	// Target the production API
	prod := godaddy.NewProductionAPI(godaddy.NewOptions(MyAPIKey, MyAPISecret))

	// Target version 1 of the production API
	prodv1 := prod.V1()

	//// You can do the same thing with the dev API as well:
	// dev := godaddy.NewDevelopmentAPI(godaddy.NewOptions(MyAPIKey, MyAPISecret))
	// devv1 := dev.V1() // etc...

	// Set our domain
	domain := prodv1.Domain(DomainToTarget)

	// Check availability
	domainDetails, err := domain.GetDetails()
	if err != nil {
		panic(err.Error())
	}

	// This could all be simplified to:
	// domainDetails, err := prod.V1().Domain(DomainToTarget).GetDetails()

	fmt.Println(domainDetails)
}
