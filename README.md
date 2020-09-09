# godaddygo

---

# Table of Contents

- [Intro](#intro)
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [API Structure](#our-api-structure)
  - [Usage](#usage)
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

### Usage 

You can either connect to the production API or the development API by running either:
 - `godaddygo.ConnectProduction(..)`
 - `godaddygo.ConnectDevelopment(...)`

We create the default client behind the scenes for you, which allows you to get to the "meat and potatoes" as fast as possible.

```golang
// Options for client
k := "api_key"
s := "api_secret"
// Skip creating client, let us create the default client
// behind the scenes
// You can now get to the core API endpoints in fewer 
// lines of code
api := godaddygo.ConnectProduction(k, s)
// or for OTE (development)...
// api := godaddygo.ConnectDevelopment(k, s)

// Use `api` here!
```

### Default Client:

With that being said you can access the default client and pass it into `endpoints.Connect(...)`

```go
package main

import (
	gdgClient "github.com/oze4/godaddygo/pkg/client"
	gdgEndpoints "github.com/oze4/godaddygo/pkg/endpoints"
)

func main() {
	// Options for client
	k := "api_key"
	s := "api_secret"
	// See here for more on GoDaddy production vs development (OTE) API's
	// https://developer.godaddy.com/getstarted
	targetProductionAPI := true

	// Create default client
	client := gdgClient.Default(k, s, targetProductionAPI)

	// Connect our client to endpoints
	api := gdgEndpoints.Connect(client)

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

### Custom Client

```go
package main

import (
	gdgEndpoints "github.com/oze4/godaddygo/pkg/endpoints"
)

func main() {
	myCustomClient := &myClient{
		key:    "api_key",
		secret: "api_secret",
		isprod: true,
	}

	api := gdgEndpoints.Connect(myCustomClient)

	// Use `api` here!

	//
	// The rest is the same as using the default client
	//
}

// As long as your client satisfies `client.Interface`
// You can use it to connect to the `endpoints` Gateway
type myClient struct {
	key    string
	secret string
	isprod bool
	// ...your custom stuff
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
