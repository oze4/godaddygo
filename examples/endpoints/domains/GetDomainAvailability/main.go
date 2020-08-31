package main

import (
	"fmt"

	godaddy "github.com/oze4/godaddygo"
)

func main() {
	MyAPIKey := "-"                // Make sure to supply appropriate key [prod or dev]
	MyAPISecret := "-"             // ...same here
	DomainToTarget := "google.com" // Domain to check availability for

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
	domAvailability, err := domain.IsAvailable()
	// Outputs a struct like:
	//     type Available struct {
	//         Available  bool   `json:"available,omitempty"`
	//         Currency   string `json:"currency,omitempty"`
	//         Definitive bool   `json:"definitive,omitempty"`
	//         Domain     string `json:"domain,omitempty"`
	//         Period     int    `json:"period,omitempty"`
	//         Price      int    `json:"price,omitempty"`
	//     }
	if err != nil {
		panic(err.Error())
	}

	// Turn availability into legible string
	isavailStr := "No."
	if domAvailability.Available == true {
		isavailStr = "Yes."
	}

	fmt.Printf("Is %s available? %s\n", DomainToTarget, isavailStr)
}
