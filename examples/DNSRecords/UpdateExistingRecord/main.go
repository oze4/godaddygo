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
	targetDomain := "-"
	// Connect to production Gateway
	api, _ := godaddygo.NewProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy Gateway
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Target `records` for this domain
	recs := domain.Records()
	// Update existing record
	newrecord := godaddygo.Record{
		Data: "1.1.0.0",
	}
	existingRecordName := "-"
	existingRecordType := godaddygo.RecordTypeA
	if err := recs.ReplaceByTypeAndName(context.Background(), existingRecordType, existingRecordName, newrecord); err != nil {
		fmt.Printf("error in TestRecordReplaceByTypeAndName : %s\n", err)
	}
	fmt.Println("success!")
}
