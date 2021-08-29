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
	// Connect to production V1 Gateway
	api, err := godaddygo.NewProduction(prodKey, prodSecret)
	if err != nil {
		panic(err.Error())
	}
	prodv1 := api.V1()
	// Get all domains you own
	mydoms, err := prodv1.ListDomains(context.Background())
	if err != nil {
		panic(err.Error())
	}

	// Loop through them
	for _, d := range mydoms {
		fmt.Println(d.Domain + " " + d.Status)
	}
}
