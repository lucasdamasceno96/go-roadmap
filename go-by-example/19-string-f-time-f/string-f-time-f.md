# Guia Didático de Strings, Templates, e Manipulação de Tempo em Go

Este guia cobre conceitos fundamentais de Go, incluindo **Funções de Strings**, **Formatação de Strings**, **Templates de Texto**, **Expressões Regulares**, **JSON**, **XML**, **Tempo**, **Epoch**, e **Formatação/Parsing de Tempo**.

---

## String Functions (Funções de Strings)

Go oferece diversas funções para manipular strings no pacote `strings`.

### Principais Funções:

- **`strings.Contains`**: Verifica se uma substring está presente.
- **`strings.ToUpper`** e **`strings.ToLower`**: Convertem para maiúsculas/minúsculas.
- **`strings.Trim`**: Remove espaços ou caracteres específicos do início/fim.
- **`strings.Split`**: Divide uma string em um slice.
- **`strings.Join`**: Junta elementos de um slice em uma única string.

#### Exemplo:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "  Golang é incrível!  "

    fmt.Println(strings.TrimSpace(text))
    fmt.Println(strings.ToUpper(text))
    fmt.Println(strings.Contains(text, "Golang"))

    words := strings.Split(text, " ")
    fmt.Println("Palavras:", words)

    fmt.Println(strings.Join(words, "-"))
}
```

---

## String Formatting (Formatação de Strings)

A formatação de strings em Go é feita usando o pacote `fmt`.

### Principais Métodos:

- **`fmt.Sprintf`**: Retorna uma string formatada.
- **`fmt.Printf`**: Imprime uma string formatada no console.
- **Placeholders Comuns**:
  - `%s`: String
  - `%d`: Inteiro
  - `%f`: Float
  - `%v`: Valor genérico

#### Exemplo:

```go
package main

import "fmt"

func main() {
    name := "Maria"
    age := 30

    formatted := fmt.Sprintf("%s tem %d anos.", name, age)
    fmt.Println(formatted)
}
```

---

## Text Templates (Templates de Texto)

Templates permitem criar textos dinâmicos substituindo variáveis em um modelo predefinido. Use o pacote `text/template`.

#### Exemplo:

```go
package main

import (
    "os"
    "text/template"
)

func main() {
    tmpl := "Olá, {{.Name}}! Você tem {{.Age}} anos."

    data := struct {
        Name string
        Age  int
    }{
        Name: "Carlos",
        Age: 25,
    }

    t, _ := template.New("example").Parse(tmpl)
    t.Execute(os.Stdout, data)
}
```

---

## Regular Expressions (Expressões Regulares)

Expressões regulares são úteis para busca e manipulação avançada de strings. Use o pacote `regexp`.

#### Exemplo:

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`[a-zA-Z]+`)
    text := "123 Golang 456"

    words := re.FindAllString(text, -1)
    fmt.Println("Palavras encontradas:", words)
}
```

---

## JSON

O pacote `encoding/json` é usado para manipular JSON.

### Exemplo de Codificação:

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    data := map[string]interface{}{
        "name": "Joana",
        "age": 28,
    }

    jsonData, _ := json.Marshal(data)
    fmt.Println(string(jsonData))
}
```

### Exemplo de Decodificação:

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    jsonData := `{"name":"Joana","age":28}`

    var data map[string]interface{}
    json.Unmarshal([]byte(jsonData), &data)

    fmt.Println(data)
}
```

---

## XML

Manipule XML com o pacote `encoding/xml`.

### Exemplo de Codificação:

```go
package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
}

func main() {
    p := Person{Name: "Carlos", Age: 40}

    xmlData, _ := xml.MarshalIndent(p, "", "  ")
    fmt.Println(string(xmlData))
}
```

---

## Time (Manipulação de Tempo)

O pacote `time` fornece ferramentas para lidar com datas e horários.

### Obtendo a Data/Hora Atual:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("Agora:", now)
}
```

---

## Epoch (Timestamp Unix)

Epoch é o número de segundos desde 1º de janeiro de 1970 (UTC).

### Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println("Timestamp Unix:", now.Unix())
}
```

---

## Time Formatting / Parsing

Use métodos como `Format` e `Parse` para manipular formatos de tempo.

### Formatando Tempo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now.Format("2006-01-02 15:04:05"))
}
```

### Parsing de Tempo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    layout := "2006-01-02 15:04:05"
    t, _ := time.Parse(layout, "2024-12-21 14:30:00")
    fmt.Println("Parsed time:", t)
}
```

---

Com estes conceitos, você pode manipular strings, templates, expressões regulares e trabalhar com tempo, JSON e XML em Go com confiança!
