package main

import "fmt"

func main() {
	fmt.Println(isPrime(11))
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}
