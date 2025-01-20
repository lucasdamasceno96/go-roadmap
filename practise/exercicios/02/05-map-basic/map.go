package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Floyd M."] = 1
	m["Ray L."] = 2
	m["Muhammed A."] = 3
	for chave, valor := range m {
		fmt.Printf("Chave: %s, Valor: %d\n", chave, valor)
	}

}
