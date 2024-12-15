# Functions em Golang

As **functions** em Golang são blocos de código reutilizáveis que permitem realizar operações específicas. Elas são uma parte fundamental do paradigma de programação funcional e organizam o código de forma modular e legível.

---

## 1. Declarando uma Função

A sintaxe básica para declarar uma função é:

```go
func nomeDaFuncao(parametros) retorno {
    // Corpo da função
}
```

### Exemplo Simples

```go
func saudacao() {
    fmt.Println("Olá, mundo!")
}

func main() {
    saudacao() // Chamando a função
}
```

---

## 2. Funções com Parâmetros

Funções podem receber um ou mais parâmetros.

```go
func soma(a int, b int) int {
    return a + b
}

func main() {
    resultado := soma(3, 4)
    fmt.Println("Resultado:", resultado)
}
```

- **Nota:** Quando os parâmetros têm o mesmo tipo, você pode simplificar a declaração:

```go
func soma(a, b int) int {
    return a + b
}
```

---

## 3. Funções com Múltiplos Retornos

Golang permite que funções retornem múltiplos valores.

```go
func dividir(a, b int) (int, int) {
    quociente := a / b
    resto := a % b
    return quociente, resto
}

func main() {
    q, r := dividir(7, 3)
    fmt.Println("Quociente:", q, "Resto:", r)
}
```

---

## 4. Funções Nomeadas e Anônimas

### Funções Nomeadas

Funções nomeadas são definidas com um identificador e podem ser chamadas diretamente.

```go
func saudacao() {
    fmt.Println("Olá!")
}
```

### Funções Anônimas

Funções anônimas são definidas sem um nome e geralmente atribuídas a variáveis ou usadas diretamente.

```go
func main() {
    mensagem := func(nome string) string {
        return "Olá, " + nome
    }
    fmt.Println(mensagem("Gopher"))
}
```

---

## 5. Funções como Parâmetros

Funções podem ser passadas como parâmetros para outras funções.

```go
func executarOperacao(a, b int, operacao func(int, int) int) int {
    return operacao(a, b)
}

func soma(a, b int) int {
    return a + b
}

func main() {
    resultado := executarOperacao(5, 3, soma)
    fmt.Println("Resultado:", resultado)
}
```

---

## 6. Closures

Closures são funções que "lembram" o contexto em que foram criadas.

```go
func contador() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    incrementar := contador()
    fmt.Println(incrementar()) // Saída: 1
    fmt.Println(incrementar()) // Saída: 2
}
```

---

## 7. Funções Variádicas

Funções variádicas aceitam um número variável de argumentos.

```go
func somaTudo(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

func main() {
    fmt.Println(somaTudo(1, 2, 3, 4)) // Saída: 10
}
```

---

## 8. Funções Recursivas

Uma função recursiva é aquela que chama a si mesma até que uma condição de parada seja atendida. Este recurso é útil para resolver problemas que podem ser divididos em subproblemas menores, como cálculo de fatorial ou Fibonacci.

### Exemplo: Fatorial

```go
func fatorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * fatorial(n-1)
}

func main() {
    fmt.Println("Fatorial de 5:", fatorial(5)) // Saída: 120
}
```

### Exemplo: Sequência de Fibonacci

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
    // Saída: 0 1 1 2 3 5 8 13 21 34
}
```

---

## 9. Exemplo Completo

```go
package main

import "fmt"

// Função para calcular a área de um retângulo
func calcularArea(largura, altura float64) float64 {
    return largura * altura
}

// Função para imprimir uma mensagem formatada
func imprimirMensagem(mensagem string) {
    fmt.Println("Mensagem:", mensagem)
}

// Função recursiva para somar números até n
func somaRecursiva(n int) int {
    if n == 0 {
        return 0
    }
    return n + somaRecursiva(n-1)
}

func main() {
    area := calcularArea(5.0, 3.5)
    imprimirMensagem(fmt.Sprintf("A área do retângulo é %.2f", area))

    soma := somaRecursiva(5)
    fmt.Println("Soma recursiva de 5:", soma) // Saída: 15
}
```

**Saída:**
```
Mensagem: A área do retângulo é 17.50
Soma recursiva de 5: 15
```

---

## Resumo

- Funções em Golang são ferramentas essenciais para modularizar e reutilizar código.
- Use parâmetros e retornos para construir funções flexíveis.
- Explore funcionalidades avançadas como múltiplos retornos, closures, funções variádicas e recursivas para resolver problemas de forma eficiente.
