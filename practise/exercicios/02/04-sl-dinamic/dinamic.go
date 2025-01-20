package main

import "fmt"

func main() {
	// Criando um slice vazio de inteiros
	var numeros []int

	// Adicionando alguns elementos
	numeros = append(numeros, 1, 2, 3)
	fmt.Printf("Capacidade inicial: %d\n", cap(numeros))

	// Adicionando mais elementos
	numeros = append(numeros, 4, 5, 6, 7, 8)
	fmt.Printf("Capacidade após adicionar mais elementos: %d\n", cap(numeros))
}
