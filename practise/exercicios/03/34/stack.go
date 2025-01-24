package main

import (
	"fmt"
)

type Stack struct {
	garage []string
}

func (s *Stack) push(car string) {
	s.garage = append(s.garage, car)
}

func (s *Stack) pop() string {
	if len(s.garage) == 0 {
		return ""
	}

	car := s.garage[len(s.garage)-1]
	s.garage = s.garage[:len(s.garage)-1]
	return car
}

func main() {
	cars := []string{"Ferrari", "Mercedes", "Red Bull", "McLaren", "Aston Martin"}
	stack := Stack{}

	// Adiciona carros à pilha
	for _, car := range cars {
		stack.push(car)
		fmt.Println("Garagem após push:", stack.garage)
	}

	// Remove carros da pilha
	for range cars { // Não precisamos do valor de `car`, apenas do tamanho da lista
		removedCar := stack.pop()
		fmt.Printf("Carro removido: %s | Garagem após pop: %v\n", removedCar, stack.garage)
	}
}
