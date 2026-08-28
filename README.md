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


### Create a Shipment

```go
shipment := &shippinglabel.Shipment{
	Carrier: &shippinglabel.Carrier{
		Code:    shippinglabel.CarrierDHL,
		Product: "V01PAK",
	},
	Sender: &shippinglabel.Address{
		Company:      "Shippinglabel",
		Street:       "Maarweg",
		StreetNumber: "62",
		PostalCode:   "50933",
		City:         "Köln",
		Country:      "DE",
	},
	Receiver: &shippinglabel.Address{
		FirstName:    "Erika",
		LastName:     "Mustermann",
		Street:       "Musterweg",
		StreetNumber: "1",
		PostalCode:   "10115",
		City:         "Berlin",
		Country:      "DE",
	},
	Parcels: []*shippinglabel.Parcel{{
		Weight: 2.5,
	}},
}

created, err := api.CreateShipment(ctx, shipment)
// Handle error

// created.Label holds the shipping label, base64 encoded
```

### FedEx

FedEx is the one carrier that cannot be used the moment its credentials are stored. Shippinglabel is
a FedEx Integrator and creates labels on the customer's own account, so FedEx requires that account
to be registered once beforehand.

Save the carrier first, then run the registration. Its first step reports which second factor FedEx
offers for the account.

```go
_, err := api.CreateCarrier(ctx, &shippinglabel.Carrier{
	Code:    shippinglabel.CarrierFedEx,
	Product: "FEDEX_INTERNATIONAL_PRIORITY",
	// Both an image type and a label stock, joined by a colon
	LabelFormat: "PDF:PAPER_4X6",
	Parameters: map[string]any{
		"accountNumber": "123456789",
	},
})
// Handle error

// The end user license agreement has to be shown to the customer in full and accepted before this
// call; the API refuses the registration without it. FedEx publishes it in English only.
opts, err := api.RegisterFedEx(ctx, &shippinglabel.FedExRegistration{
	CustomerName: "Example GmbH",
	EULAAccepted: true,
	Address: &shippinglabel.Address{
		Street:       "Maarweg",
		StreetNumber: "62",
		PostalCode:   "50933",
		City:         "Köln",
		Country:      "DE",
	},
})
// Handle error

switch {
case opts.Registered:
	// FedEx issued the credentials right away, nothing else to do

case opts.MFARequired:
	// Let the customer pick one of opts.SecureCode, or the invoice route if opts.Invoice
	err = api.GenerateFedExPin(ctx, shippinglabel.FedExPinSMS)
	// Handle error

	err = api.ConfirmFedExPin(ctx, "234560")
	// Handle error
}
```

The address has to match what FedEx holds for the account exactly, otherwise the request comes back
with `ACCOUNT.ADDRESS.MISMATCH`. Should both second factors fail, the customer can call FedEx
technical support; after that, run `RegisterFedEx` again and the response returns with
`Registered: true`.

FedEx specific parcel options travel in `Parcel.CustomFields`:

```go
parcel := &shippinglabel.Parcel{
	Weight: 2.5,
	CustomFields: map[string]any{
		shippinglabel.FedExPackagingType:   shippinglabel.FedExPackagingPak,
		shippinglabel.FedExSignatureOption: shippinglabel.FedExSignatureIndirect,
	},
}
```

Customs information is mandatory for FedEx whenever the shipment crosses a customs border.

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