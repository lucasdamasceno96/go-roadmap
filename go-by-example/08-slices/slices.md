# Slices em Golang

Os **slices** são uma estrutura fundamental em Golang para lidar com coleções de dados. Eles são mais flexíveis e poderosos que arrays, permitindo manipular sequências de elementos de maneira eficiente.

---

## 1. O que é um Slice?

Um slice é uma abstração de um array que fornece:
- Um ponteiro para o array subjacente.
- Um tamanho (`len`) que indica o número de elementos no slice.
- Uma capacidade (`cap`) que representa o número de elementos disponíveis no array subjacente a partir do primeiro elemento do slice.

---

## 2. Criando Slices

### A partir de Arrays

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Slice de arr do índice 1 ao 3
fmt.Println(slice) // Saída: [2 3 4]
```

### Usando `make`

A função `make` cria slices dinamicamente.

```go
slice := make([]int, 3, 5) // Slice com len=3 e cap=5
fmt.Println(slice)        // Saída: [0 0 0]
```

### Declaração Literal

```go
slice := []int{1, 2, 3}
fmt.Println(slice) // Saída: [1 2 3]
```

---

## 3. Propriedades de Slices

### Comprimento (`len`)

Retorna o número de elementos no slice.

```go
fmt.Println(len(slice)) // Saída: 3
```

### Capacidade (`cap`)

Retorna a capacidade máxima do slice.

```go
fmt.Println(cap(slice)) // Saída: 5 (no exemplo com `make`)
```

---

## 4. Operações Comuns

### Adicionar Elementos

A função `append` adiciona elementos a um slice. Se necessário, ela aumenta automaticamente sua capacidade.

```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)
fmt.Println(slice) // Saída: [1 2 3 4 5]
```

### Copiar Slices

A função `copy` copia elementos de um slice para outro.

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
fmt.Println(dst) // Saída: [1 2 3]
```

### Sub-slices

Você pode criar novos slices a partir de slices existentes.

```go
subSlice := slice[1:3]
fmt.Println(subSlice) // Saída: [2 3]
```

---

## 5. Diferenças entre Arrays e Slices

| **Aspecto**    | **Array**            | **Slice**         |
|-----------------|----------------------|-------------------|
| Tamanho         | Fixado na declaração | Dinâmico          |
| Mutabilidade    | Não expansível       | Expansível        |
| Alocação de Memória | Em tempo de compilação | Em tempo de execução |

---

## 6. Peculiaridades

### Compartilhamento de Dados

Slices compartilham o mesmo array subjacente, o que significa que alterações em um slice podem afetar outro que usa o mesmo array.

```go
arr := [5]int{1, 2, 3, 4, 5}
slice1 := arr[1:4]
slice2 := arr[2:5]
slice1[0] = 10
fmt.Println(slice2) // Saída: [10 4 5]
```

### Expansão

Quando a capacidade de um slice é excedida, ele cria um novo array subjacente com maior capacidade.

```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5, 6) // Capacidade aumenta
fmt.Println(slice)
```

---

## 7. Exemplo Completo

```go
package main

import "fmt"

func main() {
    // Criando um slice
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice inicial:", nums)

    // Adicionando elementos
    nums = append(nums, 6, 7)
    fmt.Println("Após append:", nums)

    // Criando um sub-slice
    sub := nums[2:5]
    fmt.Println("Sub-slice:", sub)

    // Copiando slices
    copySlice := make([]int, len(nums))
    copy(copySlice, nums)
    fmt.Println("Cópia do slice:", copySlice)

    // Alterando o sub-slice
    sub[0] = 100
    fmt.Println("Slice original após alteração no sub-slice:", nums)
}
```

**Saída:**
```
Slice inicial: [1 2 3 4 5]
Após append: [1 2 3 4 5 6 7]
Sub-slice: [3 4 5]
Cópia do slice: [1 2 3 4 5 6 7]
Slice original após alteração no sub-slice: [1 2 100 4 5 6 7]
```

---

## Resumo

- **Slices** são poderosos e flexíveis para lidar com coleções dinâmicas em Go.
- Eles compartilham o mesmo array subjacente e podem crescer dinamicamente com `append`.
- Use `len` e `cap` para monitorar o comprimento e a capacidade.
- Lembre-se de que mudanças em slices podem afetar o array ou outros slices que compartilham o mesmo array subjacente.
