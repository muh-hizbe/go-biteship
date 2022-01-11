package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//biteshipApp := biteship.New("biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLWRldiIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTYzMzk1ODk0N30.FpL4PHEikpRgNTtjXYMD85MmK35rCc0WX79n06794T4")
	biteshipApp := biteship.New("biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLXBrZyIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTY0MTc3OTY0Nn0.LA2Opjs1wNTHeSLDAZpD3W9CqMoMfZvAkOhvSYfIftk")
	//biteshipService := biteship.New("somesecret")

	//	GET LIST COURIER
	//resp, err := biteshipApp.GetCourier()

	//	CREATE SHIPMENT ORDER
	//var items []biteship.ProductItem
	//items = append(items, biteship.ProductItem{
	//	Id:          "5db7ee67382e185bd6a14608",
	//	Name:        "Black L",
	//	Image:       "",
	//	Description: "White Shirt",
	//	Value:       165000,
	//	Quantity:    1,
	//	Height:      10,
	//	Length:      10,
	//	Weight:      200,
	//	Width:       10,
	//})
	//
	//req := biteship.CreateOrderRequestParam{
	//	ShipperContactName:  "Amir",
	//	ShipperContactPhone: "081277882932",
	//	ShipperContactEmail: "biteship@test.com",
	//	ShipperOrganization: "Biteship Org Test",
	//	OriginContactName:   "Amir",
	//	OriginContactPhone:  "081740781720",
	//	OriginAddress:       "Plaza Senayan, Jalan Asia Afrik...",
	//	OriginNote:          "Deket pintu masuk STC",
	//	OriginPostalCode:    12440,
	//	OriginCoordinate: biteship.Coordinate{
	//		Latitude:  -6.2253114,
	//		Longitude: 106.7993735,
	//	},
	//	DestinationContactName:  "John Doe",
	//	DestinationContactPhone: "08170032123",
	//	DestinationContactEmail: "jon@test.com",
	//	DestinationAddress:      "Lebak Bulus MRT...",
	//	DestinationPostalCode:   12950,
	//	DestinationNote:         "Near the gas station",
	//	DestinationCoordinate: biteship.Coordinate{
	//		Latitude:  -6.28927,
	//		Longitude: 106.77492000000007,
	//	},
	//	DestinationCashOnDelivery:     nil,
	//	DestinationCashOnDeliveryType: nil,
	//	CourierCompany:                "jne",
	//	CourierType:                   "reg",
	//	CourierInsurance:              500000,
	//	DeliveryType:                  "now",
	//	DeliveryDate:                  "2022-01-11",
	//	DeliveryTime:                  "12:00",
	//	OrderNote:                     "Please be carefull",
	//	Metadata:                      nil,
	//	Items:                         items,
	//}
	//resp, err := biteshipApp.CreateOrder(&req)

	//	RETRIEVE ORDER
	//resp, err := biteshipApp.RetrieveOrder("61dce8889a096b081e70198d")

	//	UPDATE ORDER
	//requestUpdate := struct {
	//	OriginAddress string `json:"origin_address"`
	//}{OriginAddress: "Jalan Perubahan nomor 5"}
	//
	//resp, err := biteshipApp.UpdateOrder("61DBB6B1A4720916B2D1F576", requestUpdate)

	//	CONFIRM ORDER
	//resp, err := biteshipApp.ConfirmOrder("61DBB6B1A4720916B2D1F576")

	//	CANCEL ORDER
	//reason := struct {
	//	CancellationReason *string `json:"cancellation_reason"`
	//}{}
	//*reason.CancellationReason = "Ingin mengganti kurir"
	var reason string
	reason = "Ingin mengganti kurir"
	resp, err := biteshipApp.CancelOrder("61DBB6B1A4720916B2D1F576", reason)

	//	CHECK RATES OF COURIER
	//var items []biteship.ItemCourierRate
	//items = append(items, biteship.ItemCourierRate{
	//	Name:        "Shoes",
	//	Description: "Black colored size 45",
	//	Value:       199000,
	//	Length:      30,
	//	Width:       15,
	//	Height:      20,
	//	Weight:      200,
	//	Quantity:    2,
	//})
	//
	//req := biteship.RequestCourierRates{
	//	OriginLatitude:       -6.3031123,
	//	OriginLongitude:      106.7794934999,
	//	DestinationLatitude:  -6.2441792,
	//	DestinationLongitude: 106.783529000,
	//	Couriers:             "grab,jne,tiki",
	//	Items:                items,
	//}
	//
	//resp, err := biteshipApp.GetRatesCouriers(&req)

	if err != nil {
		log.Println("error :: ", err)
	}
	fmt.Println(resp)
}
