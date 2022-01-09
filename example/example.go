package main

import (
	"fmt"
	"github.com/muh-hizbe/biteship"
	"log"
)

func main() {
	req := biteship.RequestParam{}
	//biteship.Config.SecretKey = "somesecret"
	biteshipApp := biteship.New("biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLWRldiIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTYzMzk1ODk0N30.FpL4PHEikpRgNTtjXYMD85MmK35rCc0WX79n06794T4")
	resp, err := biteshipApp.CreateOrder(&req)
	//biteshipService := biteship.New("somesecret")

	//resp, err := client.CreateOrder(req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
}
