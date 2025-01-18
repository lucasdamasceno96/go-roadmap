package main

import "fmt"

func main() {
	var name string

	// takes input value for name
	fmt.Print("Enter your name:  ")
	fmt.Scanln(&name)

	fmt.Printf("Nice to ee you !!: %s", name)

}
