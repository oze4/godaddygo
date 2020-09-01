# godaddygo

Interact with the GoDaddy API via Golang

This library currently supports:

 - Getting domain info
 - Getting DNS records
 - Setting DNS records

*TERMINOLOGY*: When we reference endpoints, [this](https://developer.godaddy.com/doc) is what we are referring to.

As of now, the domains endpoint is the only endpoint that is safe to use (or available for that matter). We plan on slowly integrating other endpoints, but figure the majority of people interested in an SDK for an Domain/DNS provider (GoDaddy) prob want to view or change DNS records.

---

# Table of Contents

- [Installation](#fire-installation)
- [Details & Usage](#clipboard-usage)
- [Types](#nail_care-type-checking)
- [Examples](/examples)
- [Roadmap](#construction-roadmap)
  - [Endpoints](#endpoints)

---

<br />

## :fire: Installation

`go get -u github.com/oze4/godaddygo`

<br />

## :clipboard: Usage

Things should be fairly self-explanatory, but this is enough to get you started:

```golang
import (
	godaddy "github.com/oze4/godaddygo"
)

func main() {
	prodV1 := godaddy.NewProductionAPI(godaddy.NewOptions("myApiKey", "myApiSecret")).V1()
	// devV1 := godaddy.NewDevelopmentAPI(godaddy.NewOptions("apiKey", "apiSecret")).V1()
	myDomain := prodV1.Domain("somedomainyouown.com")
	dnsRecords := myDomain.Records().GetAll()

	// ...do something with records
}
```

<br />

## :nail_care: Type Checking

- :star:**For user-land-type-checking, each endpoint has it's own package located at:** `github.com/oze4/godaddygo/pkg/endpoints/<endpoint>`
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

## :construction: Roadmap

### Endpoints

Please see [here](https://developer.godaddy.com/doc) for more information on GoDaddy API endpoints

:no_entry_sign:<small>Not Available</small>
**|**
:construction:<small>In Progress</small>
**|**
:white_check_mark:<small>Finished</small>

| Endpoint      | Status          |
| ------------- | --------------- |
| Abuse         | :no_entry_sign: |
| Aftermarket   | :no_entry_sign: |
| Agreements    | :no_entry_sign: |
| Certificates  | :no_entry_sign: |
| Countries     | :no_entry_sign: |
| Domains       | :white_check_mark: Safe to get domain info and DNS records, as well as set DNS records  |
| Orders        | :no_entry_sign: |
| Shoppers      | :no_entry_sign: |
| Subscriptions | :no_entry_sign: |

<br />
<br />
<br />

[mattoestreich.com](https://mattoestreich.com)
