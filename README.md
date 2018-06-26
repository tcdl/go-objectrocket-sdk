# Golang SDK for ObjectRocket API
This is Go client for ObjectRocket API https://app.objectrocket.com/. Which provide create of new instances in ObjectRocket via API.

## Installation
Once you have a working Go installation locally, you can grab this package with the following command:
```
go get github.com/tcdl/go-objectrocket-sdk
```

### Creating a client
The client will be required to perform all actions against the API.
```go
package main

import (
    "github.com/tcdl/go-objectrocket-sdk"
)

func main() {
    // To request API Token you should make GET request to https://sjc-api.objectrocket.com/v2/tokens/
    // Token should be renew each 24 hours
    client := objectrocket.NewClient("API-TOKEN", nil)
}
```

### Add new instance
```go
// First initiate details for your Instance
newInstance := objectrocket.InstanceDetail{
    Name: "TEST_NAME",
    Plan: 1,
    Service: "mongodb",
    Type: "mongodb_singledb",
    Version: "3.4.10",
    Zone: "US-East-IAD2"}
result, err := client.Instance.Create_Instance(newInstance)
```
