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
- [Extended Example](#extended-example)
- [Features](#features)

---

# Intro

- See [here](#features) for more on which features this package currently supports
- Whenever we reference endpoints, [this is what we are referring to](https://developer.godaddy.com/doc)

<br />

<small>Pull requests welcome! We would like to eventually support each GoDaddy Gateway endpoint, not just domain/DNS related tasks</small>

# Installation

- `go get -u github.com/oze4/godaddygo`
- See [here](https://developer.godaddy.com/) for more on how to obtain an Gateway key and Gateway secret from GoDaddy (click 'API Keys')

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

	// Target version 1 of the production API
	godaddy := api.V1()

	//
	// See `Extended Example` section below for more
	//
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

	//
	// See `Extended Example` section below for more
	//
}
```

# Extended Example

### Regardless of your client, how you actually use this package will be the same either way.

```go
/* We are continuing from within `main()`
 * ... pretend code from above is here,
 * regardless of your client */

// We now have access to "all" GoDaddy production
// version 1 gateway endpoints (via `godaddy`)

// !! the following is pseudo code !!

foo := godaddy.Domain("foo.com")
bar := godaddy.Domain("bar.com")
// ...more domains...

// Get domain details
foo.GetDetails(ctx)
bar.GetDetails(ctxtwo)

// Anything you can do with `foo`
// you can do with `bar`

// Domain records
fooRecs := foo.Records()
// Do stuff with records
fooRecs.List(ctx)
fooRecs.Add(ctx, someDNSRecord)
fooRecs.FindByType(ctx, godaddygo.RecordTypeA)

// Account related tasks

// View all domains for your account
godaddy.ListDomains(ctx)
// Check availability for domain you don't own
godaddy.CheckAvailability(ctx, "fizz.buzz")
// Purchase domain (this hasn't been tested - it should use the card you have on file)
// I'm not sure what happens when you don't have a card on file =/ lmk
godaddy.Purchase(ctx, myNewDomain)

// etc...
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
