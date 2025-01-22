package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
}

func fibonacci(number int) int {
	if number < 2 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
}