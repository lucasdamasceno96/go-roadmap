package main

import "fmt"

func main() {
	type sneaker struct {
		model string
		price float64
		quant int
	}

	// Método que calcula o total
	func (s sneaker) total() float64 {
		return s.price * float64(s.quant)
	}

	// Instanciando a struct
	var sn1 sneaker
	sn1.model = "Air Max 90"
	sn1.price = 120.00
	sn1.quant = 10

	// Exibindo o total usando o método
	fmt.Println(sn1)
	fmt.Printf("Total: $%.2f\n", sn1.total())
}
