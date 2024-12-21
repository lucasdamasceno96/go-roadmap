# Introdução às Goroutines em Go

As goroutines são uma das principais funcionalidades da linguagem Go, permitindo a execução de tarefas concorrentes de forma eficiente. Elas são leves, rápidas e simples de usar, sendo ideais para programas que precisam lidar com múltiplas tarefas simultâneas.

---

## O que são Goroutines?
Uma goroutine é uma função que é executada de forma concorrente com outras goroutines dentro do mesmo programa. Elas são mais leves do que as threads do sistema operacional, pois:

- São gerenciadas pelo runtime do Go.
- Consomem menos memória (cerca de 2 KB por goroutine ao iniciar).
- Escalam melhor, permitindo que milhares (ou até milhões) de goroutines sejam executadas ao mesmo tempo.

---

## Criando Goroutines

Para criar uma goroutine, basta usar a palavra-chave `go` antes de uma chamada de função. Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func printMessage(message string) {
    for i := 0; i < 5; i++ {
        fmt.Println(message)
        time.Sleep(100 * time.Millisecond) // Simula uma tarefa demorada
    }
}

func main() {
    go printMessage("Goroutine 1") // Executa esta função como uma goroutine
    go printMessage("Goroutine 2")

    // O programa principal também é uma goroutine
    printMessage("Main Goroutine")
}
```

### Saída esperada
A saída pode variar, pois as goroutines são executadas de forma concorrente:

```
Main Goroutine
Goroutine 1
Goroutine 2
Main Goroutine
Goroutine 1
...
```

---

## Comunicação entre Goroutines

As goroutines podem se comunicar e sincronizar usando canais (channels). Um canal permite enviar e receber valores entre goroutines.

### Exemplo de uso de canais

```go
package main

import "fmt"

func sum(a int, b int, result chan int) {
    result <- a + b // Envia o resultado para o canal
}

func main() {
    result := make(chan int) // Cria um canal

    go sum(3, 5, result) // Executa a função como uma goroutine

    fmt.Println("Resultado:", <-result) // Recebe o valor do canal
}
```

---

## Tratando Concorrência com `sync.WaitGroup`

Se você precisa garantir que todas as goroutines terminem antes de prosseguir, pode usar `sync.WaitGroup`:

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Marca a goroutine como concluída
    fmt.Printf("Worker %d começando\n", id)
    fmt.Printf("Worker %d terminou\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Adiciona uma goroutine ao WaitGroup
        go worker(i, &wg)
    }

    wg.Wait() // Espera todas as goroutines terminarem
    fmt.Println("Todos os workers concluíram.")
}
```

---

## Goroutines e Threads

### Diferenças entre Goroutines e Threads:
1. **Leveza**: Goroutines têm um overhead muito menor do que threads.
2. **Multiplexação**: O runtime do Go gerencia goroutines em um pool fixo de threads do SO.
3. **Custo de criação**: Criar goroutines é significativamente mais barato do que criar threads.

---

## Conclusão
As goroutines são uma ferramenta poderosa para escrever programas concorrentes em Go. Combinadas com canais e estruturas como `sync.WaitGroup`, elas tornam a programação concorrente mais simples, eficiente e legível. No entanto, é importante evitar problemas como condições de corrida, garantindo que o acesso a recursos compartilhados seja devidamente sincronizado.
