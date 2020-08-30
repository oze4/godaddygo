# godaddygo

Interact with the GoDaddy API via Golang

- [Installation](#fire-installation)
- [Details & Usage](#clipboard-details-and-usage)
  - [Type-checking](#type-checking)
- [Examples](/examples)
- [Roadmap](#construction-roadmap)
  - [Endpoints](#endpoints)

---

## :fire: Installation

`go get -u github.com/oze4/godaddygo`

## :clipboard: Details and Usage

- To import:

```golang
import (
  // ...
  "github.com/oze4/godaddygo"
)
```

- _The main package, (`github.com/oze4/godaddygo`) provides everything you need to interact with the GoDaddy API_.
- **However**, if you would like to type-check using types which are returned from `godaddygo`, you will need to use a specific package for "that" endpoint
- We mirrored endpoints found in the [GoDaddy docs](https://developer.godaddy.com/doc)
- :star:**For user-land-type-checking, each endpoint has it's own package located at:** `github.com/oze4/godaddygo/pkg/<endpoint>`

### Type-checking

If you wanted to type-check using the data type for domain details, you would do:

```golang
// ...

import (
  // ...
  "github.com/oze4/godaddygo/pkg/domains"
)

// SomeFunc says 'leave me alone, vscode' 😉
func SomeFunc(dd domains.DomainDetails) { // <-- type-check using `DomainDetails`
  // ...
}
```

## :construction: Roadmap

**:exclamation:THIS LIBRARY SHOULD \*\*NOT BE USED AT ALL\*\* RIGHT NOW!**:exclamation:as it is under heavy construction

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
| Domains       | :construction:  |
| Orders        | :no_entry_sign: |
| Shoppers      | :no_entry_sign: |
| Subscriptions | :no_entry_sign: |
