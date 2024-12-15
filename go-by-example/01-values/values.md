# Golang Values

Este repositório lista exemplos de como usar valores e expressões comuns no Golang. Isso inclui variáveis, constantes, operadores e funções nativas.

---

## Conteúdo

1. [Constantes e Variáveis](#constantes-e-variáveis)
2. [Tipos de Dados](#tipos-de-dados)
3. [Operadores Matemáticos](#operadores-matemáticos)
4. [Operadores Lógicos](#operadores-lógicos)
5. [Funções Comuns](#funções-comuns)
6. [Exemplo Completo](#exemplo-completo)

---

## Constantes e Variáveis

### Constantes

```go
const pi = 3.14
const hello = "Olá, Mundo!"
const truth = true
```

- As constantes têm valores fixos e não podem ser alteradas durante a execução.

### Variáveis

```go
var name string = "Golang"
var age int = 10
var isCool bool = true
```

- As variáveis podem ser declaradas usando `var` ou a notação curta:

```go
name := "Golang"
age := 10
isCool := true
```

## Tipos de Dados

Golang possui tipos primitivos e compostos. Aqui estão os principais:

### Tipos Primitivos

| Tipo     | Exemplo             |
|----------|---------------------|
| `int`    | `var x int = 42`    |
| `float64`| `var y float64 = 3.14` |
| `bool`   | `var z bool = true` |
| `string` | `var s string = "Hello"` |

### Tipos Compostos

| Tipo         | Exemplo                            |
|--------------|------------------------------------|
| Array        | `var arr [3]int = [3]int{1, 2, 3}` |
| Slice        | `var slice []int = []int{1, 2, 3}` |
| Map          | `var m map[string]int`            |
| Struct       | `type User struct { Name string }` |
| Pointer      | `var ptr *int = &x`               |

## Operadores Matemáticos

```go
// Soma
result := 3 + 2

// Subtração
result := 5 - 3

// Multiplicação
result := 2 * 3

// Divisão
result := 10 / 2

// Módulo
result := 10 % 3
```

## Operadores Lógicos

```go
// E lógico
fmt.Println(true && false) // false

// OU lógico
fmt.Println(true || false) // true

// NÃO lógico
fmt.Println(!true) // false
```

## Funções Comuns

### Funções nativas

#### Manipulação de Strings

```go
fmt.Println(len("Golang"))            // Comprimento da string
fmt.Println("Hello" + " " + "World") // Concatenação
```

#### Conversão de Tipos

```go
x := 42
f := float64(x) // Conversão para float64
s := string(x)  // Conversão para string
```

#### Entrada e Saída

```go
fmt.Println("Hello, World!")
fmt.Printf("Valor: %d\n", 42)
```

## Exemplo Completo

Aqui está um exemplo que reúne todos os conceitos:

```go
package main

import "fmt"

func main() {
    const pi = 3.14
    name := "Golang"
    age := 10
    isCool := true

    fmt.Println("Nome:", name)
    fmt.Println("Idade:", age)
    fmt.Println("Legal:", isCool)

    result := 10 + 20
    fmt.Println("Resultado da soma:", result)

    fmt.Println("Comparação lógica:", true && false)
    fmt.Println("Comprimento da string:", len("Golang"))
    
    fmt.Printf("Pi: %.2f\n", pi)
}
```

---

Explore e experimente com os exemplos para entender os valores e operadores do Golang!
