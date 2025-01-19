package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 5
	a[1] = 10
	var sum int
	for i := 0; i < 2; i++ {
		sum = sum + a[i]
	}
	fmt.Println("Sum: ", sum)
}
