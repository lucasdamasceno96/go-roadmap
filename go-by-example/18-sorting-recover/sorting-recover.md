# Guia Didático: Sorting, Panic, Defer e Recover em Go

Este guia vai explicar conceitos importantes de Go de forma simples e prática: **Sorting**, **Sorting by Functions**, **Panic**, **Defer**, e **Recover**.

---

## Sorting (Ordenação)

Ordenação é o processo de organizar elementos em uma sequência específica, como ordem crescente ou decrescente.

### Ordenando com `sort.Ints` e `sort.Strings`
Go fornece funções prontas no pacote `sort` para ordenar slices de inteiros e strings.

#### Exemplo:

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    nums := []int{5, 3, 8, 1, 2}
    sort.Ints(nums)
    fmt.Println("Números ordenados:", nums)

    names := []string{"João", "Ana", "Carlos", "Beatriz"}
    sort.Strings(names)
    fmt.Println("Nomes ordenados:", names)
}
```

**Saída:**
```
Números ordenados: [1 2 3 5 8]
Nomes ordenados: [Ana Beatriz Carlos João]
```

---

## Sorting by Functions (Ordenação Personalizada)

Quando você quer ordenar um slice baseado em critérios específicos, pode usar `sort.Slice`.

### Exemplo:

```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {Name: "Alice", Age: 25},
        {Name: "Bob", Age: 30},
        {Name: "Charlie", Age: 20},
    }

    // Ordenando por idade
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println("Pessoas ordenadas por idade:", people)
}
```

**Saída:**
```
Pessoas ordenadas por idade: [{Charlie 20} {Alice 25} {Bob 30}]
```

---

## Panic

**Panic** é usado para lidar com erros críticos no programa que impedem sua continuação. Quando ocorre um `panic`, o programa termina, mas você pode recuperar (ou interceptar) esse erro usando `defer` e `recover`.

### Exemplo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Antes do panic")

    panic("Algo deu muito errado!")

    // Este código nunca será executado
    fmt.Println("Depois do panic")
}
```

**Saída:**
```
Antes do panic
panic: Algo deu muito errado!
```

---

## Defer

**Defer** adia a execução de uma função até que a função que a contém termine. É muito útil para garantir que recursos sejam liberados corretamente.

### Exemplo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Início")

    defer fmt.Println("Fim")

    fmt.Println("Meio")
}
```

**Saída:**
```
Início
Meio
Fim
```

### Usando `defer` para fechar arquivos

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        panic(err)
    }

    // Certifica-se de fechar o arquivo
    defer file.Close()

    fmt.Fprintln(file, "Escrevendo no arquivo")
}
```

---

## Recover

**Recover** é usado para capturar um erro que causaria um `panic` e evitar que o programa termine abruptamente.

### Exemplo:

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado do erro:", r)
        }
    }()

    fmt.Println("Antes do panic")
    panic("Erro crítico!")
    fmt.Println("Depois do panic") // Nunca será executado
}
```

**Saída:**
```
Antes do panic
Recuperado do erro: Erro crítico!
```

---

## Combinação de Defer, Panic e Recover

Combinando `defer`, `panic` e `recover`, você pode criar programas resilientes.

### Exemplo Completo:

```go
package main

import "fmt"

func safeDivision(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado do erro:", r)
        }
    }()

    if b == 0 {
        panic("Divisão por zero")
    }

    return a / b
}

func main() {
    fmt.Println("Resultado:", safeDivision(10, 2))
    fmt.Println("Resultado:", safeDivision(10, 0))
    fmt.Println("Programa continua...")
}
```

**Saída:**
```
Resultado: 5
Recuperado do erro: Divisão por zero
Programa continua...
```

---

Com esses conceitos e exemplos, você agora entende como lidar com ordenação, erros e controle de fluxo em Go. Pratique esses conceitos para criar programas robustos e eficientes!
