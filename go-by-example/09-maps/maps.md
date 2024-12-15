# Maps em Golang

Os **maps** em Golang são coleções que armazenam pares chave-valor. Eles são úteis para situações em que é necessário associar um valor a uma chave específica.

---

## 1. Criando Maps

### Usando `make`

A função `make` é a forma mais comum de criar um map.

```go
m := make(map[string]int)
```

### Declaração Literal

Você também pode criar maps e inicializá-los com valores.

```go
m := map[string]int{
    "a": 1,
    "b": 2,
}
```

---

## 2. Operações Básicas

### Adicionar e Atualizar Valores

Para adicionar ou atualizar um valor no map, basta usar a chave correspondente.

```go
m := make(map[string]int)
m["chave"] = 42
fmt.Println(m["chave"]) // Saída: 42
```

### Ler Valores

Use a chave para acessar o valor.

```go
valor := m["chave"]
fmt.Println(valor)
```

### Remover Valores

A função `delete` remove uma entrada do map.

```go
delete(m, "chave")
```

### Verificar Existência de Chave

Você pode usar a segunda variável de retorno para verificar se uma chave existe.

```go
valor, ok := m["chave"]
if ok {
    fmt.Println("Chave encontrada, valor:", valor)
} else {
    fmt.Println("Chave não encontrada")
}
```

---

## 3. Iterando sobre Maps

Use um laço `for range` para iterar sobre os pares chave-valor.

```go
m := map[string]int{
    "a": 1,
    "b": 2,
}

for chave, valor := range m {
    fmt.Printf("Chave: %s, Valor: %d\n", chave, valor)
}
```

**Saída:**
```
Chave: a, Valor: 1
Chave: b, Valor: 2
```

---

## 4. Características dos Maps

### Tamanho Dinâmico

Os maps crescem dinamicamente conforme novos elementos são adicionados.

### Ordem Não Garantida

A ordem de iteração sobre um map não é garantida.

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}
for k, v := range m {
    fmt.Println(k, v) // A ordem pode variar
}
```

### Chaves Válidas

As chaves devem ser de tipos que possam ser comparados com `==`. Exemplos de tipos válidos:
- `string`
- `int`
- `bool`

Tipos como slices, maps e funções não podem ser usados como chaves.

---

## 5. Exemplo Completo

```go
package main

import "fmt"

func main() {
    // Criando um map
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    // Adicionando um valor
    m["d"] = 4

    // Atualizando um valor
    m["a"] = 10

    // Removendo um valor
    delete(m, "b")

    // Iterando sobre o map
    for chave, valor := range m {
        fmt.Printf("Chave: %s, Valor: %d\n", chave, valor)
    }

    // Verificando se uma chave existe
    if val, ok := m["c"]; ok {
        fmt.Println("A chave 'c' existe com valor:", val)
    } else {
        fmt.Println("A chave 'c' não existe")
    }
}
```

**Saída:** (a ordem pode variar)
```
Chave: a, Valor: 10
Chave: c, Valor: 3
Chave: d, Valor: 4
A chave 'c' existe com valor: 3
```

---

## Resumo

- Os **maps** são estruturas flexíveis para armazenar pares chave-valor.
- Use `make` para criar maps ou declarações literais para inicializá-los.
- Funções como `delete` e verificações com `ok` ajudam a gerenciar os dados de maneira eficiente.
- A iteração sobre maps não garante ordem, e as chaves devem ser de tipos comparáveis.

Explore os maps para resolver problemas que exigem associações rápidas e dinâmicas entre chaves e valores!
