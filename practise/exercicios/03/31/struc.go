package main

import "fmt"

func main() {
	type sneaker struct {
		model string
		price float64
		quant int
	}

	var sn1 sneaker
	sn1.model = "Air Max 90"
	sn1.price = 120.00
	sn1.quant = 10
	fmt.Println(sn1)
}
