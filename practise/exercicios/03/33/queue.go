package main

import (
	"fmt"
	"time"
)

type Queue struct { // Nome do tipo deve ser exportado (começar com maiúscula)
	drivers []string
}

func (q *Queue) push(driver string) {
	q.drivers = append(q.drivers, driver)
}

func (q *Queue) pop() string {
	if len(q.drivers) == 0 {
		return ""
	}

	driver := q.drivers[0]
	q.drivers = q.drivers[1:]
	return driver
}

func main() {
	q := Queue{} // Criando uma instância da fila

	q.push("Hamilton") // Usando o método push para adicionar pilotos
	q.push("Verstappen")
	q.push("Leclerc")

	fmt.Println("Iniciando PitStop...")
	time.Sleep(1 * time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println("Entrando Piloto:", q.pop())
		time.Sleep(1 * time.Second)
	}
}
