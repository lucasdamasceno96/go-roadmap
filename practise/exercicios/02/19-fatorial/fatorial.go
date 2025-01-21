package main

import "fmt"

func main() {
	fmt.Println(fatorial(5))
}

func fatorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * fatorial(number-1) // executando a recursão
}
