# Estruturas, Métodos, Interfaces e Genéricos em Go

Go é uma linguagem poderosa que suporta conceitos fundamentais como estruturas (structs), métodos, interfaces e genéricos. Esses conceitos formam a base para organizar e reutilizar código de maneira eficiente.

---

## Estruturas (Structs)

### O que são Structs?
Uma `struct` é uma coleção de campos que agrupam dados relacionados. Em Go, structs são amplamente utilizadas para modelar dados e construir aplicações.

### Criando Structs

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p)
}
```

### Acessando e Modificando Campos

```go
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name) // Alice

p.Age = 31
fmt.Println(p.Age) // 31
```

### Ponteiros com Structs

```go
func updateAge(person *Person, newAge int) {
    person.Age = newAge
}

p := &Person{Name: "Bob", Age: 40}
updateAge(p, 45)
fmt.Println(p.Age) // 45
```

---

## Métodos

### O que são Métodos?
Métodos são funções associadas a um tipo específico (geralmente structs). Eles são definidos usando um *receiver* (receptor).

### Criando Métodos

```go
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

p := Person{Name: "Alice", Age: 30}
p.Greet()
```

### Métodos com Receivers Ponteiros
Use ponteiros como receivers para modificar o valor do struct:

```go
func (p *Person) HaveBirthday() {
    p.Age++
}

p.HaveBirthday()
fmt.Println(p.Age) // 31
```

---

## Interfaces

### O que são Interfaces?
Interfaces definem um conjunto de métodos que um tipo deve implementar. Elas permitem polimorfismo em Go.

### Definindo Interfaces

```go
type Greeter interface {
    Greet()
}

type Person struct {
    Name string
}

func (p Person) Greet() {
    fmt.Println("Hello, I am", p.Name)
}
```

### Usando Interfaces

```go
func SayHello(g Greeter) {
    g.Greet()
}

p := Person{Name: "Alice"}
SayHello(p)
```

### Interface Vazia
A interface vazia (`interface{}`) pode armazenar qualquer tipo, mas deve ser usada com cuidado.

```go
var value interface{} = 42
fmt.Println(value)
```

---

## Genéricos

### O que são Genéricos?
Genéricos permitem criar funções e tipos que podem operar em diferentes tipos de dados, aumentando a reutilização do código.

### Funções Genéricas

```go
func PrintSlice[T any](slice []T) {
    for _, v := range slice {
        fmt.Println(v)
    }
}

PrintSlice([]int{1, 2, 3})
PrintSlice([]string{"A", "B", "C"})
```

### Tipos Genéricos

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

stack := Stack[int]{}
stack.Push(10)
stack.Push(20)
fmt.Println(stack.Pop()) // 20
```

### Restrições em Genéricos
Você pode restringir os tipos permitidos em genéricos usando interfaces:

```go
type Adder[T interface{ int | float64 }] struct {
    A, B T
}

func (a Adder[T]) Sum() T {
    return a.A + a.B
}

adder := Adder[int]{A: 10, B: 20}
fmt.Println(adder.Sum()) // 30
```

---

## Conclusão

- **Structs**: Modelam dados de forma estruturada.
- **Métodos**: Adicionam comportamento a tipos específicos.
- **Interfaces**: Facilitam abstração e polimorfismo.
- **Genéricos**: Permitem reutilização de código com diferentes tipos de dados.

Esses conceitos juntos tornam Go uma linguagem poderosa e flexível para desenvolvimento de software eficiente e organizado.
