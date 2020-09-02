This library currently supports:

 - Getting domain info
 - Checking domain availability
 - Getting DNS records
 - Setting DNS records

Whenever we reference endpoints, [this is what we are referring to](https://developer.godaddy.com/doc)

Pull requests welcome! We plan on slowly integrating each GoDaddy endpoint

---

# Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Types](#type-checking)
- [Examples](https://github.com/oze4/godaddygo/tree/master/examples)
- [Roadmap](#roadmap)
  - [Endpoints](#endpoints)

---

<br />

## Installation

`go get -u github.com/oze4/godaddygo`

<br />

## Usage

This library is rather "self-documenting", but this is enough to get you started:

```golang
import (
	godaddy "github.com/oze4/godaddygo"
)

func main() {
	prodV1 := godaddy.NewProductionAPI(godaddy.NewOptions("myApiKey", "myApiSecret")).V1()
	
	// Can also target the development API (see the GoDaddy docs for more details)
	// devV1 := godaddy.NewDevelopmentAPI(godaddy.NewOptions("myApiKey", "myApiSecret")).V1()
	
	// 9 times out of 10 you will need to own the domain for your query
	// You can check for domain availability without owning the domain, for exmaple
	myDomain := prodV1.Domain("mydomain.com")
	myDomainRecords := myDomain.Records()
	myDomainZone := myDomainRecords.GetAll()

	// ...do something with `myDomainZone`
}
```

<br />

## Type Checking

- **For user-land-type-checking, each endpoint has it's own package located at:** `github.com/oze4/godaddygo/pkg/endpoints/<endpoint>`
- _The main package, (`github.com/oze4/godaddygo`) provides everything you need to interact with the GoDaddy API_.
- **However**, if you would like to type-check using types which are returned from `godaddygo`, you will need to use a specific package for "that" endpoint
- We mirrored endpoints found in the [GoDaddy docs](https://developer.godaddy.com/doc)

If you wanted to type-check using the data type for domain details, you would do:

```golang
package xyz

import (
  "github.com/oze4/godaddygo/pkg/endpoints/domains"
)

func SomeFunc(dd domains.DomainDetails) { // <-- type-check using `DomainDetails`
  // ...
}
```

<br />

## Roadmap

### Endpoints

Please see [here](https://developer.godaddy.com/doc) for more information on GoDaddy API endpoints

| Endpoint      | Status          |
| ------------- | --------------- |
| Abuse         | - |
| Aftermarket   | - |
| Agreements    | - |
| Certificates  | - |
| Countries     | - |
| Domains       | Safe to get domain info and DNS records, as well as set DNS records  |
| Orders        | - |
| Shoppers      | - |
| Subscriptions | - |

<br />
<br />
<br />

[mattoestreich.com](https://mattoestreich.com)
