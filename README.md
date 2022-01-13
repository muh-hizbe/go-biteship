# Biteship API Go
This library is unofficial Biteship for golang.

- [Documentation](#documentation)
- [Installation](#installation)
- [Usage](#usage)
  - [Couriers](#courier)
  - [Rates](#rates)
  - [Order](#order)
  - [Tracking](#tracking)

##  Documentation
For the API documentation, please get by contact Biteship or check at [biteship.com](https://biteship.com).

## Installation
```shell
go get -u github.com/muh-hizbe/biteship
```
```go
import "github.com/muh-hizbe/biteship"
```

##  Usage
You must have secretKey for using with this library, you can get it from dashboard Biteship.
Check [here](https://github.com/muh-hizbe/go-biteship/blob/master/example/example.go) for detail of example usage
### Courier
####    Get Couriers
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)

	//	GET LIST COURIER
	resp, err := biteshipApp.GetCourier()
	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```
### Rates
####    Check Rates Couriers
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)

	//	CHECK RATES OF COURIER
	var items []biteship.ItemCourierRate
	items = append(items, biteship.ItemCourierRate{
		Name:        "Shoes",
		Description: "Black colored size 45",
		Value:       199000,
		Length:      30,
		Width:       15,
		Height:      20,
		Weight:      200,
		Quantity:    2,
	})

	req := biteship.RequestCourierRates{
		//OriginLatitude:       -6.3031123,
		//OriginLongitude:      106.7794934999,
		DestinationLatitude:  -6.2441792,
		DestinationLongitude: 106.783529000,
		OriginPostalCode:     12440,
		//DestinationPostalCode: 12240,
		//Couriers: "jne,tiki",
		Items: items,
	}

	resp, err := biteshipApp.GetRatesCouriers(&req)
	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```
### Order
####    Create an Order
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)
	
	// CREATE SHIPMENT ORDER
	var items []biteship.ProductItem
	items = append(items, biteship.ProductItem{
		Id:          "5db7ee67382e185bd6a14608",
		Name:        "Black L",
		Image:       "",
		Description: "White Shirt",
		Value:       165000,
		Quantity:    1,
		Height:      0,
		Length:      0,
		Width:       0,
		Weight:      0,
	})

	req := biteship.CreateOrderRequestParam{
		ShipperContactName:  "Amir",
		ShipperContactPhone: "081277882932",
		ShipperContactEmail: "biteship@test.com",
		ShipperOrganization: "Biteship Org Test",
		OriginContactName:   "Amir",
		OriginContactPhone:  "081740781720",
		OriginAddress:       "Plaza Senayan, Jalan Asia Afrik...",
		OriginNote:          "Deket pintu masuk STC",
		OriginPostalCode:    12440,
		OriginCoordinate: biteship.Coordinate{
			Latitude:  -6.2253114,
			Longitude: 106.7993735,
		},
		DestinationContactName:  "John Doe",
		DestinationContactPhone: "08170032123",
		DestinationContactEmail: "jon@test.com",
		DestinationAddress:      "Lebak Bulus MRT...",
		DestinationPostalCode:   12950,
		DestinationNote:         "Near the gas station",
		DestinationCoordinate: biteship.Coordinate{
			Latitude:  -6.28927,
			Longitude: 106.77492000000007,
		},
		DestinationCashOnDelivery:     nil,
		DestinationCashOnDeliveryType: nil,
		CourierCompany:                "jne",
		CourierType:                   "reg",
		CourierInsurance:              500000,
		DeliveryType:                  "now",
		DeliveryDate:                  "2022-01-11",
		DeliveryTime:                  "12:00",
		OrderNote:                     "Please be carefull",
		Metadata:                      nil,
		Items:                         items,
	}
	
	resp, err := biteshipApp.CreateOrder(&req)
	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```
####    Retrieve an Order
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)
	
	//	RETRIEVE AN ORDER
	resp, err := biteshipApp.RetrieveOrder("61dce8889a096b081e70198d")
	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```
####    Update an Order
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)
	
	//	UPDATE ORDER
	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}
	
	resp, err := biteshipApp.UpdateOrder("61DBB6B1A4720916B2D1F576", requestUpdate)
	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```
####    Confirm an Order
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)
	
	//  CONFIRM ORDER
	resp, err := biteshipApp.ConfirmOrder("61DBB6B1A4720916B2D1F576")

	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```
####    Cancel an Order
```go
package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//	CHANGE WITH YOUR SECRET KEY
	secretKey := "replace_with_your_secret_key"
	biteshipApp := biteship.New(secretKey)
	
	//	CANCEL ORDER
	var reason string
	reason = "Ingin mengganti kurir"
	resp, err := biteshipApp.CancelOrder("61DBB6B1A4720916B2D1F576", reason)
	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
```