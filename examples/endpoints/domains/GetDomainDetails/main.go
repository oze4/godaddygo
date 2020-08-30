package main

import (
	"fmt"
	"github.com/oze4/godaddygo"
)

func main() {
    MyAPIKey := "-"     // Make sure to supply appropriate key [prod or dev]
    MyAPISecret := "-"  // ...same here
    DomainToTarget := "somedomainyouown.com" // *** YOU MUST OWN THIS DOMAIN TO GET DETAILS ***

    clientOptions := godaddygo.Options{
        APIKey: MyAPIKey, 
        APISecret: MyAPISecret, 
   }
   client := godaddygo.NewClient(clientOptions)

   // Target the production API
   prod := client.NewProduction()

   // Target version 1 of the production API
   prodv1 := prod.V1()

   //// You can do the same thing with the dev API as well:
   // dev := client.NewDevelopment()
   // devv1 := dev.V1()

   // Set our domain
   domain := prodv1.Domain(DomainToTarget)

   // Check availability
   domainDetails, err := domain.GetDetails()
   if err != nil {
       panic(err.Error())
   }

   fmt.Println(domainDetails)
}