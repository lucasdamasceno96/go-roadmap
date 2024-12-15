package main

import "fmt"

var x int = 11 // fora de funcoes tem que utilizar o var

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" // := somente aceito dentro de funcoes
	fmt.Println(f)

	fmt.Println(x)
}
