package main

import (
	"context"
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "// your prod key"
	prodSecret := "// your prod secret"
	targetDomain := "dom.com"
	// Connect to production Gateway
	api, _ := godaddygo.NewProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy Gateway
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Target `records` for this domain
	records := domain.Records()
	// Update existing record
	recordToUpdate := godaddygo.Record{
		Name: "echo",
		Type: godaddygo.RecordTypeA,
		Data: "1.2.3.4",
		TTL: 360,
	}
	if err := records.Update(context.Background(), []godaddygo.Record{recordToUpdate}); err != nil {
		panic("Unable to update record : " + err.Error())
	}
	fmt.Println("Success! Updated record")
}
