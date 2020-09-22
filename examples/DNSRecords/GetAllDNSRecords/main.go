package main

import (
	"fmt"

	"github.com/oze4/godaddygo"
)

func main() {
	// Your info...
	prodKey := "-"
	prodSecret := "-"
	targetDomain := "dom.com"
	// Connect to production Gateway
	api := godaddygo.ConnectProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy Gateway
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Target `records` for this domain
	records := domain.Records()
	// Get all records
	dnsrecords, err := records.GetAll() // -> *[]DNSRecord
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(dnsrecords)
}
