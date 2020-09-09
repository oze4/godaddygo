# godaddygo

---

# Table of Contents

- [Intro](#intro)
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [API Structure](#our-api-structure)
  - [Default Client](#default-client)
  - [Custom Client](#custom-client)
- [Examples](https://github.com/oze4/godaddygo/tree/master/examples)
- [Roadmap](#roadmap)
  - [Endpoints](#endpoints)

---

## Intro

Hello and welcome! Thanks for checking us out :smile:

**This library currently supports**:

- Getting domain info
- Checking domain availability
- Getting DNS records
- Setting DNS records

Whenever we reference endpoints, [this is what we are referring to](https://developer.godaddy.com/doc)

Pull requests welcome! We plan on slowly integrating each GoDaddy endpoint

## Installation

`go get -u github.com/oze4/godaddygo`

## Getting Started

The `godaddygo` package is essentially for convenience. You have the ability to create your own client which you can pass to `endpoints.Connect(<your_client>)`

### Our API Structure

Consider the following endpoint, which allows you to add a DNS record to a domain.

![screenshot_from_godaddy_docs](https://i.imgur.com/tN2IveY.png)

Programmatically, this would look like:

```golang
// Simplified
api.V1().Domain("dom.com").Records().Add(newDNSRecord)
```


### Default Client:

```go
// Recommended way
package main

import (
    "github.com/oze4/godaddygo"
)

func main() {
    // Options for client
    k := "api_key"
    s := "api_secret"
    // See here for more on GoDaddy production vs development (OTE) API's
    // https://developer.godaddy.com/getstarted
    targetProductionAPI := true

    // Create new client
    client := godaddygo.NewClient(targetProductionAPI, k, s)

    // Connect our client to endpoints
    api := godaddygo.Connect(client)

    //
	// Use `api` here!
	//
    // For example:
    prodv1 := api.V1()
    // Target specific domain
    mydomain := prodv1.Domain("mydomain.com")
    // Get all DNS records for target domain
    records, err := mydomain.Records().GetAll()

    // ...
}
```

You can also circumvent having to specify whether or not to target the production API.

```golang
// Simplified

// Options for client
k := "api_key"
s := "api_secret"
// Create new client
client := godaddygo.ConnectProduction(k, s)
// or for OTE (development)...
// client := godaddygo.ConnectDevelopment(k, s)

// etc...
```

### Custom Client

```go
package main

import (
    "github.com/oze4/godaddygo/pkg/endpoints"
)

func main() {
    // Instead of doing `godaddy := godaddygo.Connect(client)`, which
	// just wraps around `endpoints.Connect`, you would do:
	myCustomClient := &myClient{
		key: "api_key",
		secret: "api_secret",
		isprod: true,
	}

    api := endpoints.Connect(myCustomClient)

    // Use `api` here!

    //
    // The rest is the same as using the default client
    //
}

// As long as your client satisfies `client.Interface`
// You can use it to connect to the `endpoints` Gateway
type myClient struct {
    key string
    secret string
    isprod bool
    // ...
}

func (c *myClient) APIKey() string {
    return c.key
}

func (c *myClient) APISecret() string {
    return c.secret
}

func (c *myClient) IsProduction() string {
    return c.isprod
}
```

## Roadmap

### Endpoints

Please see [here](https://developer.godaddy.com/doc) for more information on GoDaddy API endpoints

- [ ] Abuse
- [ ] Aftermarket
- [ ] Agreements
- [ ] Certificates
- [ ] Countries
- [x] Domains
  - [x] DNS Records
- [ ] Orders
- [ ] Shoppers
- [ ] Subscriptions

<br />
<br />
<br />

[mattoestreich.com](https://mattoestreich.com)
