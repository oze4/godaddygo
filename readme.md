# [godaddygo](https://oze4.github.io/godaddygo/)

[Check us out on `pkg.go.dev`](https://pkg.go.dev/github.com/oze4/godaddygo?tab=doc)

---

# Table of Contents

- [godaddygo](#godaddygo)
- [Table of Contents](#table-of-contents)
- [Intro](#intro)
- [Installation](#installation)
- [Usage](#usage)
	- [Basic Usage](#basic-usage)
	- [Custom Client](#custom-client)
- [Features](#features)

---

# Intro

 - See [here](#features) for more on which features this package currently supports
 - Whenever we reference endpoints, [this is what we are referring to](https://developer.godaddygo.com/doc)

<br /> 

<small>Pull requests welcome! We would like to eventually support each GoDaddy Gateway endpoint, not just domain/DNS related tasks</small>

# Installation

 - `go get -u github.com/oze4/godaddygo`
 - See [here](https://developer.godaddygo.com/) for more on how to obtain an Gateway key and Gateway secret from GoDaddy (click 'Gateway Keys')

# Usage

## Basic Usage

Bare minimum what you need to get started (aka how you will typically use this package):

```golang
package main

import (
	"github.com/oze4/godaddygo"
)

func main() {
	key := "<your_key>"
	secret := "<your_secret>"

	// Target production GoDaddy API
	// 99% of the time this is what you are looking for
	api, err := godaddygo.NewProduction(key, secret)
	if err != nil {
		panic(err.Error())
	}
	
	// Target version 1 of production API
	godaddy := api.V1() 

	// Now have access to all GoDaddy production V1 Gateway endpoints (via `godaddy`)
	dom := godaddy.Domain("foo.bar")
	// Get domain details
	dom.GetDetails(ctx)
	// Domain records
	domRecs := myDomain.Records()
	domRecs.List(ctx)
	domRecs.Add(ctx, someDNSRecord)
	domRecs.FindByType(ctx, godaddygo.RecordTypeA)
	// View all domains for your account
	godaddy.ListDomains(ctx)
	// Check availability for domain
	godaddy.CheckAvailability(ctx, "baz.bar")
	// Purchase domain
	godaddy.Purchase(ctx, myNewDomain)

	// etc...
}
```

## Custom Client

```go
package main

import (
	"net/http"

	"github.com/oze4/godaddygo"
)

func main() {
	key := "<your_key>"
	secret := "<your_secret>"
	// Target production API (godaddygo.APIDevEnv | godaddygo.APIProdEnv)
	target := godaddygo.APIProdEnv 

	// Build new config
	myConfig := godaddygo.NewConfig(key, secret, target)
	// Build custom client
	myClient := &http.Client{}

	// Establish "connection" with API
	api, err := godaddygo.WithClient(myClient, myConfig)
	if err != nil {
		panic(err.Error())
	}

	// Target version 1 of the production API
	godaddy := api.V1()

	// Now have access to all GoDaddy production V1 Gateway endpoints (via `godaddy`)
	dom := godaddy.Domain("foo.bar")
	// Get domain details
	dom.GetDetails(ctx)
	// Domain records
	domRecs := myDomain.Records()
	domRecs.List(ctx)
	domRecs.Add(ctx, someDNSRecord)
	domRecs.FindByType(ctx, godaddygo.RecordTypeA)
	// View all domains for your account
	godaddy.ListDomains(ctx)
	// Check availability for domain
	godaddy.CheckAvailability(ctx, "baz.bar")
	// Purchase domain
	godaddy.Purchase(ctx, myNewDomain)

	// etc...
}
```

# Features

Please see [here](https://developer.godaddygo.com/doc) for more information on GoDaddy Gateway endpoints

- [ ] Abuse
- [ ] Aftermarket
- [ ] Agreements
- [ ] Certificates
- [ ] Countries
- Domains
  - [x] Check domain availability
  - [x] Get all DNS records
  - [x] Get all DNS records of specific type
  - [x] Get specific DNS record
  - [x] Set DNS record(s)
  - [x] Add/create DNS record(s)
  - [x] Delete/remove DNS record(s)
  - [x] Purchase domain
  - [x] Purchase privacy for domain
  - [x] Remove privacy for domain
- [ ] Orders
- [ ] Shoppers
- [ ] Subscriptions

<br />
<br />
<br />

---

[mattoestreich.com](https://mattoestreich.com)

---
