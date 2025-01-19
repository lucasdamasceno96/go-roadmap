package main

import "fmt"

func main() {
	var name string
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Print("Enter your name:  ")
	fmt.Scanln(&name)

	// Atualize a variável names com o valor retornado por addName
	names = addName(name, names)

	fmt.Println(names, len(names)) // Agora o novo nome será exibido no slice
}

func addName(name string, names []string) []string {
	// Adicione o novo nome ao slice
	names = append(names, name)
	return names // Retorne o slice atualizado
}
