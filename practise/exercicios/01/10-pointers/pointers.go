package main

import "fmt"

func main() {
	x := 10
	p := &x

	*p = 20 // Altera o valor de x através do ponteiro

	fmt.Println("Valor de x:", x) // 20
}
