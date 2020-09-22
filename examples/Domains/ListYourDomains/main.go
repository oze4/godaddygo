package main

import (
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "-"
	prodSecret := "-"
	// Connect to production V1 Gateway
	api := godaddygo.ConnectProduction(prodKey, prodSecret).V1()
	// Get all domains you own
	mydoms, err := api.Domains().My()
	if err != nil {
		panic(err.Error())
	}

	// Loop through them
	for _, d := range *mydoms {
		fmt.Println(d.Domain + " " + d.Status)
	}
}
