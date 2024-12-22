# Guia Didático: Random Numbers, Number Parsing, URL Parsing, SHA256 Hashes e Base64 Encoding em Go

Este guia aborda conceitos essenciais de Go relacionados a números aleatórios, parsing de números e URLs, geração de hashes SHA256 e codificação Base64.

---

## Random Numbers (Números Aleatórios)

Em Go, o pacote `math/rand` é usado para gerar números aleatórios.

### Gerando Números Aleatórios

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // Inicializa a semente com a hora atual

    fmt.Println("Número aleatório entre 0 e 100:", rand.Intn(100))
    fmt.Println("Número aleatório de ponto flutuante:", rand.Float64())
}
```

---

## Number Parsing (Parsing de Números)

O pacote `strconv` é usado para converter strings em números e vice-versa.

### Exemplo de Conversão

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "42"
    num, err := strconv.Atoi(str) // Converte string para inteiro
    if err != nil {
        fmt.Println("Erro:", err)
    }
    fmt.Println("Número convertido:", num)

    floatStr := "3.14"
    fnum, err := strconv.ParseFloat(floatStr, 64) // Converte string para float64
    if err != nil {
        fmt.Println("Erro:", err)
    }
    fmt.Println("Float convertido:", fnum)
}
```

---

## URL Parsing (Parsing de URLs)

O pacote `net/url` é usado para analisar e manipular URLs.

### Exemplo:

```go
package main

import (
    "fmt"
    "net/url"
)

func main() {
    rawURL := "https://example.com/path?query=value"

    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        fmt.Println("Erro ao analisar URL:", err)
    }

    fmt.Println("Esquema:", parsedURL.Scheme)
    fmt.Println("Host:", parsedURL.Host)
    fmt.Println("Caminho:", parsedURL.Path)
    fmt.Println("Query:", parsedURL.RawQuery)
}
```

---

## SHA256 Hashes (Hashes SHA256)

Para calcular hashes SHA256, use o pacote `crypto/sha256`.

### Exemplo:

```go
package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    data := "Go é incrível!"

    hash := sha256.Sum256([]byte(data))
    fmt.Printf("Hash SHA256: %x\n", hash)
}
```

---

## Base64 Encoding (Codificação Base64)

A codificação Base64 é realizada usando o pacote `encoding/base64`.

### Exemplo:

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    data := "Aprendendo Go!"

    encoded := base64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println("Codificado em Base64:", encoded)

    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        fmt.Println("Erro na decodificação:", err)
    }
    fmt.Println("Decodificado:", string(decoded))
}
```

---

Com esses exemplos, você agora entende como manipular números aleatórios, realizar parsing de números e URLs, gerar hashes SHA256 e usar codificação Base64 em Go. Pratique e explore esses conceitos!
