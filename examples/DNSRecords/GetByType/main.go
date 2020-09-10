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
	// Connect to production API
	api := godaddygo.ConnectProduction(prodKey, prodSecret)
	// Target version 1 of the production GoDaddy API
	prodv1 := api.V1()
	// Set our domain
	domain := prodv1.Domain(targetDomain)
	// Target `records` for this domain
	records := domain.Records()
	// Get all `A` records
	dnsrecords, err := records.GetByType("A") // -> *[]DNSRecord
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(dnsrecords)
}
