package main

import "fmt"

// Função para obter o nome do boxeador com base no ranking
func getBoxerName(ranking int, boxers map[string]int) string {
        for name, rank := range boxers {
                if rank == ranking {
                        return name
                }
        }
        return "Boxeador não encontrado."
}

func main() {
        m := make(map[string]int)
        m["Floyd M."] = 1
        m["Ray L."] = 2
        m["Muhammed A."] = 3

        fmt.Println("Enter a number between 1 and 3 and see who is the best boxer: ")
        var boxernumber int
        fmt.Scanln(&boxernumber)

        boxerName := getBoxerName(boxernumber, m) 
        if boxerName != "Boxeador não encontrado." {
                fmt.Printf("Fighter: %s, Ranked at: %d\n", boxerName, boxernumber)
        } else {
                fmt.Println("Ranking inválido.")
        }
}package main

import "fmt"

// Função para obter o nome do boxeador com base no ranking
func getBoxerName(ranking int, boxers map[string]int) string {
        for name, rank := range boxers {
                if rank == ranking {
                        return name
                }
        }
        return "Boxeador não encontrado."
}

func main() {
        m := make(map[string]int)
        m["Floyd M."] = 1
        m["Ray L."] = 2
        m["Muhammed A."] = 3

        fmt.Println("Enter a number between 1 and 3 and see who is the best boxer: ")
        var boxernumber int
        fmt.Scanln(&boxernumber)

        boxerName := getBoxerName(boxernumber, m) 
        if boxerName != "Boxeador não encontrado." {
                fmt.Printf("Fighter: %s, Ranked at: %d\n", boxerName, boxernumber)
        } else {
                fmt.Println("Ranking inválido.")
        }
}