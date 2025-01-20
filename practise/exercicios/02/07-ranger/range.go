package main

import "fmt"

func main() {
	capitais := map[string]string{
		"Brasil":         "Brasília",
		"Estados Unidos": "Washington D.C.",
		"França":         "Paris",
	}

	// Iterando sobre o mapa e imprimindo as chaves e valores
	for pais, capital := range capitais {
		fmt.Printf("País: %s, Capital: %s\n", pais, capital)
	}
}
