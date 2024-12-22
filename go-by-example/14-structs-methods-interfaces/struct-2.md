# Struct Embedding em Go

**Struct Embedding** é um recurso do Go que permite que uma struct inclua outra como um campo embutido. Isso facilita a reutilização de código e é uma forma simples de implementar composição, substituindo a herança que outras linguagens oferecem.

Quando você faz o embedding de uma struct, os campos e métodos da struct embutida são promovidos para a struct que os contém, permitindo que você os acesse diretamente, como se fossem definidos na própria struct que faz o embedding.

---

## Exemplo de Struct Embedding

### Código Básico
```go
package main

import "fmt"

// Definindo uma struct básica
type Person struct {
    Name string
    Age  int
}

// Definindo outra struct que faz embedding da struct Person
type Employee struct {
    Person // Embedding
    Position string
}

func main() {
    emp := Employee{
        Person: Person{Name: "Alice", Age: 30},
        Position: "Developer",
    }

    // Acessando campos da struct embutida diretamente
    fmt.Println(emp.Name)       // Alice
    fmt.Println(emp.Age)        // 30
    fmt.Println(emp.Position)   // Developer
}
```

No exemplo acima:
- `Person` foi embutida em `Employee`.
- Os campos `Name` e `Age` da struct `Person` estão diretamente acessíveis através de `Employee`.

---

## Métodos em Struct Embedding

Quando uma struct embutida possui métodos, esses métodos também são promovidos para a struct que faz o embedding.

### Exemplo com Métodos

```go
package main

import "fmt"

// Struct básica com um método
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Struct com embedding
type Employee struct {
    Person
    Position string
}

func main() {
    emp := Employee{
        Person: Person{Name: "Bob", Age: 40},
        Position: "Manager",
    }

    // Chamando o método da struct embutida
    emp.Greet()
}
```

**Saída:**
```
Hello, my name is Bob and I am 40 years old.
```

---

## Benefícios do Struct Embedding

1. **Reutilização de Código**: Struct Embedding permite reaproveitar campos e métodos de structs existentes.
2. **Composição**: Go incentiva a composição em vez de herança, e embedding é uma das maneiras de implementar isso.
3. **Acesso Simplificado**: Campos e métodos da struct embutida podem ser acessados diretamente sem precisar de qualificação explícita.

---

## Conflitos de Nome em Struct Embedding

Se a struct que faz o embedding e a struct embutida tiverem campos ou métodos com o mesmo nome, será necessário desambiguar explicitamente.

### Exemplo de Conflito

```go
package main

import "fmt"

type Person struct {
    Name string
}

type Employee struct {
    Person
    Name string // Conflito de campo
}

func main() {
    emp := Employee{
        Person: Person{Name: "Alice"},
        Name:   "Bob",
    }

    fmt.Println(emp.Name)        // Bob (campo de Employee)
    fmt.Println(emp.Person.Name) // Alice (campo de Person)
}
```

---

## Enumeradores (Enums) em Go

Go não possui suporte nativo para enums como algumas outras linguagens, mas você pode simular enums usando constantes e o tipo `iota`.

### O que é `iota`?
`iota` é um identificador predefinido no Go que é usado para gerar números sequenciais dentro de blocos de constantes.

### Exemplo de Enum Simulado

```go
package main

import "fmt"

// Definindo um enum para estados de pedidos
type OrderStatus int

const (
    Pending OrderStatus = iota
    Processing
    Shipped
    Delivered
)

func (s OrderStatus) String() string {
    return []string{"Pending", "Processing", "Shipped", "Delivered"}[s]
}

func main() {
    status := Pending
    fmt.Println(status) // Pending

    status = Shipped
    fmt.Println(status) // Shipped
}
```

### Benefícios de Usar Enums com `iota`

1. **Legibilidade**: Facilita a leitura e o entendimento do código.
2. **Controle de Valores**: Garante que as variáveis podem assumir apenas valores definidos no enum.
3. **Flexibilidade**: Pode ser combinado com métodos para adicionar comportamentos específicos aos valores.

### Limitando Valores com Enums

Você pode usar enums para validar estados:

```go
func IsValidStatus(status OrderStatus) bool {
    switch status {
    case Pending, Processing, Shipped, Delivered:
        return true
    default:
        return false
    }
}

func main() {
    fmt.Println(IsValidStatus(3)) // true
    fmt.Println(IsValidStatus(5)) // false
}
```

---

## Conclusão

Struct Embedding é uma técnica poderosa e idiomática no Go para compor tipos e reaproveitar código. Ao mesmo tempo, enums simulados com `iota` oferecem uma maneira eficiente de representar conjuntos de valores nomeados, promovendo organização e legibilidade no código.
