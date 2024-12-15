# Uso do `for` em Golang

O laço `for` em Golang é a única estrutura de repetição da linguagem, mas é extremamente versátil e pode ser usado em diversos cenários.

---

## 1. Sintaxe Básica

A forma mais comum de usar o `for` é semelhante a outras linguagens:

```go
for inicialização; condição; incremento {
    // Código a ser executado
}
```

### Exemplo

```go
for i := 0; i < 5; i++ {
    fmt.Println("Valor de i:", i)
}
```

**Saída:**
```
Valor de i: 0
Valor de i: 1
Valor de i: 2
Valor de i: 3
Valor de i: 4
```

---

## 2. `for` como um `while`

Se você omitir a inicialização e o incremento, o `for` funciona como um `while`.

```go
x := 0
for x < 5 {
    fmt.Println("Valor de x:", x)
    x++
}
```

**Saída:**
```
Valor de x: 0
Valor de x: 1
Valor de x: 2
Valor de x: 3
Valor de x: 4
```

---

## 3. Laço Infinito

O `for` pode ser usado sem nenhum argumento, criando um laço infinito. Você deve incluir uma instrução para sair do laço (como `break`).

```go
for {
    fmt.Println("Laço infinito")
    break // Sai do laço
}
```

---

## 4. `for` com `range`

O `for range` é usado para iterar sobre coleções como arrays, slices, maps, strings e canais.

### Iterando sobre um slice

```go
nums := []int{1, 2, 3, 4}
for index, value := range nums {
    fmt.Printf("Índice: %d, Valor: %d\n", index, value)
}
```

**Saída:**
```
Índice: 0, Valor: 1
Índice: 1, Valor: 2
Índice: 2, Valor: 3
Índice: 3, Valor: 4
```

### Iterando sobre uma string

```go
str := "Golang"
for index, char := range str {
    fmt.Printf("Índice: %d, Caractere: %c\n", index, char)
}
```

**Saída:**
```
Índice: 0, Caractere: G
Índice: 1, Caractere: o
Índice: 2, Caractere: l
Índice: 3, Caractere: a
Índice: 4, Caractere: n
Índice: 5, Caractere: g
```

### Iterando sobre um map

```go
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("Chave: %s, Valor: %d\n", key, value)
}
```

---

## 5. Controle do Fluxo

- **`break`**: Sai do laço imediatamente.
- **`continue`**: Pula para a próxima iteração.

### Exemplo com `break` e `continue`

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Sai do laço quando i é 5
    }
    if i%2 == 0 {
        continue // Pula números pares
    }
    fmt.Println(i)
}
```

**Saída:**
```
1
3
```

---

## Resumo

- O `for` é usado para todas as formas de iteração em Go.
- Pode atuar como um `while`, `foreach` ou até criar laços infinitos.
- É combinado com palavras-chave como `break` e `continue` para controle de fluxo.
- O `for range` é a maneira mais idiomática de iterar sobre coleções em Go.

Explore essas variações para resolver problemas de iteração de maneira eficaz!
