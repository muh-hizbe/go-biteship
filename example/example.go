package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	//req := biteship.CreateOrderRequestParam{}
	//biteship.Config.SecretKey = "somesecret"
	//biteshipApp := biteship.New("biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLWRldiIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTYzMzk1ODk0N30.FpL4PHEikpRgNTtjXYMD85MmK35rCc0WX79n06794T4")
	biteshipApp := biteship.New("biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLXBrZyIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTY0MTc3OTY0Nn0.LA2Opjs1wNTHeSLDAZpD3W9CqMoMfZvAkOhvSYfIftk")
	//resp, err := biteshipApp.CreateOrder(&req)
	resp, err := biteshipApp.GetCourier()
	//biteshipService := biteship.New("somesecret")

	//resp, err := client.CreateOrder(req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
}
