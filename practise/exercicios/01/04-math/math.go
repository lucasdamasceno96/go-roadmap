package main

import "fmt"

func main() {
	var num1, num2 int
	var result int
	var multi int
	var div int
	var sub int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	result = num1 + num2
	multi = num1 * num2
	div = num1 / num2
	sub = num1 - num2

	fmt.Printf("Result of %d + %d is %d", num1, num2, result)
	fmt.Printf("\nResult of %d * %d is %d", num1, num2, multi)
	fmt.Printf("\nResult of %d / %d is %d", num1, num2, div)
	fmt.Printf("\nResult of %d - %d is %d", num1, num2, sub)
}
