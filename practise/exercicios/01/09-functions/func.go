package main

import "fmt"

var a int
var b int

func main() {
	a = 5
	b = 10
	fmt.Println(sum(a, b))
}

func sum(a int, b int) int {
	return a + b
}
