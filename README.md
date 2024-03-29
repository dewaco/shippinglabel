# Golang SDK for the Shippinglabel REST API

This Golang package is an SDK for the Shippinglabel REST API (https://shippinglabel.de). Documentation can be found
at https://developer.shippinglabel.de

## Usage

### Install

``` sh
go get github.com/dewaco/shippinglabel
```

### Create Client and API Context

```go
package main

import (
	"context"
	"github.com/dewaco/shippinglabel"
)

func main() {
	// Create request client
	client, err := shippinglabel.NewClient("CLIENT_ID", "CLIENT_SECRET")
	// Handle error

	// Create an access token
	ctx := context.Background()
	tk, err := client.ClientCredentials(ctx)
	// Handle error

	// Create a context
	api, err := client.APIContext(tk)
	// Handle error
	
	// Request: Get user details
	user, err := api.GetUser(ctx)
	// Handle error
	
	// Request: Create parcel
	parcel, err := api.CreateParcel(ctx, &shippinglabel.Parcel{})
	// Handle error
	
	// Request: Delete a shipment
	err = api.DeleteShipment(ctx, 1)
	// Handle error
	
	// ...
}
```

### Handle Response Error

```go
package main

import (
	"context"
	"github.com/dewaco/shippinglabel"
)

func main() {
	client, err := shippinglabel.NewClient("CLIENT_ID", "SECRET")
	ctx := context.Background()
	tk, err := client.ClientCredentials(ctx)
	api, err := client.APIContext(tk)
	
	// Send request
	_, err := api.GetUser(ctx)
	
	// Parse error
	if err != nil {
		switch err.(type) {
		case *shippinglabel.Error:
			// Is SL error
			break
		default:
			// Is an unexpected error
		}
		
		// Another way
		_, ok := err.(*shippinglabel.Error)
		if ok { 
			// Is SL error
                }
	}
}
```