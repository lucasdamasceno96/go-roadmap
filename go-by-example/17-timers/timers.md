# Concurrency e Sincronização em Go

Go oferece diversas ferramentas para lidar com tarefas concorrentes e sincronização. Aqui estão explicações detalhadas sobre **Timers**, **Tickers**, **Worker Pools**, **Wait Groups**, **Rate Limiting**, **Atomic Counters**, **Mutexes**, e **Stateful Goroutines**, com exemplos práticos.

---

## Timers

**Timers** permitem agendar uma ação para ser executada após um período de tempo.

### Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTimer(2 * time.Second)

    fmt.Println("Timer started")

    <-timer.C // Aguarda o canal do timer
    fmt.Println("Timer expired")
}
```

### Cancelando um Timer:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTimer(2 * time.Second)

    go func() {
        <-timer.C
        fmt.Println("Timer expired")
    }()

    stop := timer.Stop() // Cancela o timer
    if stop {
        fmt.Println("Timer stopped")
    }
}
```

---

## Tickers

**Tickers** disparam eventos repetidamente em intervalos regulares.

### Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
    quit := make(chan struct{})

    go func() {
        for {
            select {
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            case <-quit:
                ticker.Stop()
                return
            }
        }
    }()

    time.Sleep(5 * time.Second)
    close(quit)
}
```

---

## Worker Pool

**Worker Pools** ajudam a controlar o número de goroutines que processam tarefas simultaneamente.

### Exemplo:

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2 // Simula um trabalho
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    // Criando 3 workers
    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }

    // Enviando 5 trabalhos
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Recebendo resultados
    for r := 1; r <= 5; r++ {
        fmt.Println("Result:", <-results)
    }
}
```

---

## Wait Groups

**Wait Groups** permitem sincronizar a conclusão de múltiplas goroutines.

### Exemplo:

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Marca esta goroutine como concluída
    fmt.Printf("Worker %d started\n", id)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Adiciona uma goroutine ao grupo
        go worker(i, &wg)
    }

    wg.Wait() // Aguarda todas as goroutines terminarem
    fmt.Println("All workers done")
}
```

---

## Rate Limiting

**Rate Limiting** controla a frequência de execução de tarefas.

### Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("Processing request", req, "at", time.Now())
    }
}
```

---

## Atomic Counters

**Atomic Counters** garantem operações seguras entre goroutines usando o pacote `sync/atomic`.

### Exemplo:

```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {
    var counter int64

    for i := 0; i < 10; i++ {
        go func() {
            for {
                atomic.AddInt64(&counter, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)
    fmt.Println("Counter:", atomic.LoadInt64(&counter))
}
```

---

## Mutexes

**Mutexes** protegem seções críticas de código para evitar condições de corrida.

### Exemplo:

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    var count int

    increment := func() {
        mu.Lock()
        defer mu.Unlock()
        count++
    }

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }

    wg.Wait()
    fmt.Println("Count:", count)
}
```

---

## Stateful Goroutines

**Stateful Goroutines** mantêm estado interno isolado dentro de uma goroutine específica.

### Exemplo:

```go
package main

import "fmt"

func counter() chan int {
    ch := make(chan int)
    go func() {
        count := 0
        for {
            ch <- count
            count++
        }
    }()
    return ch
}

func main() {
    countCh := counter()

    for i := 0; i < 5; i++ {
        fmt.Println(<-countCh)
    }
}
```

---

Essas ferramentas são fundamentais para implementar soluções concorrentes e sincronizadas em Go. Elas oferecem controle fino sobre execução paralela, tornando seu código eficiente e seguro.
