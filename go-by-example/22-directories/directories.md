# Trabalhando com Arquivos em Go

Este documento aborda como trabalhar com arquivos em Go, incluindo leitura de arquivos, escrita de arquivos e filtragem de linhas. Os exemplos são baseados nos conceitos encontrados em [Go by Example](https://go-by-example.com/).

## 1. Lendo Arquivos (Reading Files)

Em Go, você pode usar o pacote `os` e o pacote `io/ioutil` (ou `os` e `bufio` para leitura linha por linha) para ler arquivos. Aqui está um exemplo simples de como ler o conteúdo de um arquivo:

### Exemplo: Lendo um Arquivo Inteiro

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    // Lendo o conteúdo de um arquivo
    content, err := ioutil.ReadFile("example.txt")
    if err != nil {
        log.Fatal(err)
    }

    // Exibindo o conteúdo do arquivo
    fmt.Println(string(content))
}
