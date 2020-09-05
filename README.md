# godaddygo

---

# Table of Contents

- [Intro](#intro)
- [Installation](#installation)
- [Usage](#usage)
  - [Default Client](#default-client)
  - [Custom Client](#custom-client)
  - [API Structure](#overall-api-structure)
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

## Usage

### Default Client:

The `godaddygo` package wraps around the `/pkg/endpoints` package.

```go
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
    godaddy := godaddygo.Connect(client)

    //
    // Use `godaddy` here!
    //

    // For example
    prodv1 := godaddy.V1() // godaddy.V2() etc..
    // Target specific domain
    mydomain := prodv1.Domain("mydomain.com")
    // Get all DNS records for target domain
    records, err := mydomain.Records().GetAll()

    // ...
}
```

### Custom Client

Again, *the `godaddygo` package wraps around the `/pkg/endpoints` package*. **This means you have the ability to write your own client**. As long as your client satisfies the [`session.Interface`](https://github.com/oze4/godaddygo/blob/master/pkg/session/interface.go#L3) interface, you can pass it to [`endpoints.NewConnection(client)`](https://github.com/oze4/godaddygo/blob/master/pkg/endpoints/connection.go#L29).

This would look something like:

```go
package main

import (
    "github.com/oze4/godaddygo/pkg/endpoints"
)

func main() {
    // Instead of doing: 
    //// godaddy := godaddygo.Connect(client)
    // Which ultimaely just wraps around `endpoints.NewConnection()`,
    // you would do:
    godaddy := endpoints.NewConnection(myclient) // pretend `myclient` satisfies `session.Interface`

    //
    // Use `godaddy` here! The rest is the same as using
    // the default client
    //

    // For example
    prodv1 := godaddy.V1() // godaddy.V2() etc..
    // Target specific domain
    mydomain := prodv1.Domain("mydomain.com")
    // Get all DNS records for target domain
    records, err := mydomain.Records().GetAll()

    // ...
}
```

### Overal API Structure

Consider the following endpoint, which allows you to add a DNS record to a domain..

![screenshot_from_godaddy_docs](https://i.imgur.com/tN2IveY.png)

Programmatically, this would look like:

```golang
// Simplified
godaddy.V1().Domain("dom.com").Records().Add(newDNSRecord)
```

## Roadmap

### Endpoints

Please see [here](https://developer.godaddy.com/doc) for more information on GoDaddy API endpoints

| Endpoint      | Status                                                              |
| ------------- | ------------------------------------------------------------------- |
| Abuse         | -                                                                   |
| Aftermarket   | -                                                                   |
| Agreements    | -                                                                   |
| Certificates  | -                                                                   |
| Countries     | -                                                                   |
| Domains       | Safe to get domain info and DNS records, as well as set DNS records |
| Orders        | -                                                                   |
| Shoppers      | -                                                                   |
| Subscriptions | -                                                                   |

<br />
<br />
<br />

[mattoestreich.com](https://mattoestreich.com)
