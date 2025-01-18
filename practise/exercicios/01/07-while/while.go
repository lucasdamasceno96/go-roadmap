package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number:  ")
	fmt.Scanln(&number)
	i := 1
	for i <= number {
		fmt.Println(i)
		i = i + 1
	}
}
