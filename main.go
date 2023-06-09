package main

import (
	"simple-payment/delivery"
)

func main() {
	delivery.NewServer("localhost:8080").Run()
}
