package main

import (
	"fmt"
	"live-coding/helpers"
	"log"
)

func main() {
	var customer1, err = helpers.GetCustomerByID("1", true, "e3cb86bd7a10483c8d2a7f7f35527036")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>>>>>Customer 1 api result >>>>>>")
	fmt.Println(customer1)
	customer2, err := helpers.GetCustomerByID("2", true, "e3cb86bd7a10483c8d2a7f7f35527036")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>>>>>Customer 2 api result >>>>>>")
	fmt.Println(customer2)
}
