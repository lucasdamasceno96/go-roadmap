# Roadmap para Estruturas de Dados e Algoritmos (DSA) em Go (Golang)

## Introdução
Este roadmap é uma guia passo a passo para aprender Estruturas de Dados e Algoritmos (DSA) utilizando Go. Inclui explicações sobre conceitos fundamentais, como notação Big O, cálculo de complexidade de algoritmos, e uma lista de algoritmos essenciais para dominar.

---

## O que estudar?

1. **Conceitos Fundamentais:**
   - Variáveis, loops, funções e arrays.
   - Estruturas de controle (if/else, switch, etc.).
   - Manejo de erros em Go.

2. **Estruturas de Dados:**
   - Arrays e Slices.
   - Listas encadeadas.
   - Pilhas e Filas.
   - Árvores e Grafos.
   - Hash Maps (ou Maps em Go).
   - Heaps.

3. **Conceitos de Algoritmos:**
   - Busca e Ordenação.
   - Recursão.
   - Algoritmos de Grafos (BFS, DFS).
   - Algoritmos de Divisão e Conquista.

---

## Notação Big O
Big O é usada para descrever o desempenho de um algoritmo em termos de:

1. **Tempo de Execução:** Quantidade de operações que o algoritmo realiza para diferentes tamanhos de entrada.
2. **Uso de Espaço:** Quantidade de memória adicional usada pelo algoritmo.

### Tipos de Complexidade:
- **O(1):** Constante - o tempo de execução é o mesmo, independente do tamanho da entrada.
- **O(n):** Linear - cresce proporcionalmente ao tamanho da entrada.
- **O(n²):** Quadrático - o tempo de execução aumenta exponencialmente com o tamanho da entrada.
- **O(log n):** Logarítmico - reduz pela metade a cada iteração.

### Exemplo de Cálculo:
#### Código:
```go
func linearSearch(arr []int, target int) bool {
    for _, num := range arr {
        if num == target {
            return true
        }
    }
    return false
}
```
#### Análise:
1. O loop percorre todos os elementos do array: **O(n)**.
2. Espaço adicional usado é constante: **O(1)**.

---

## Lista de Algoritmos Essenciais

### 1. **Ordenação**
   - **Bubble Sort:** Ordenação simples, mas ineficiente (O(n²)).
   - **Merge Sort:** Algoritmo de divisão e conquista (O(n log n)).
   - **Quick Sort:** Altamente eficiente na média (O(n log n)).

### 2. **Busca**
   - **Busca Linear:** Percorre todos os elementos (O(n)).
   - **Busca Binária:** Requer array ordenado; muito eficiente (O(log n)).

### 3. **Recursão**
   - Conceito fundamental em algoritmos como busca em profundidade e Fibonacci.

### 4. **Grafos**
   - **BFS (Busca em Largura):** Visita níveis de um grafo (O(V + E)).
   - **DFS (Busca em Profundidade):** Explora profundamente até os nós terminais (O(V + E)).

### 5. **Programação Dinâmica**
   - Resolve problemas dividindo em subproblemas menores.
   - Exemplos: Problema da Mochila, Subsequência Comum Máxima (LCS).

### 6. **Algoritmos de Caminho Mínimo**
   - **Dijkstra:** Calcula o caminho mais curto de um nó para todos os outros (O(V²) ou O(V + E log V) com Heap).
   - **Floyd-Warshall:** Calcula caminhos mais curtos entre todos os pares de nós (O(V³)).

### 7. **Divisão e Conquista**
   - **Exemplo:** Algoritmo de multiplicação de grandes números.

---

## Detalhes dos Algoritmos

### Bubble Sort:
- **Descrição:** Compara elementos adjacentes e os troca se estiverem fora de ordem.
- **Complexidade:** O(n²).

### Merge Sort:
- **Descrição:** Divide o array em partes menores e as combina de forma ordenada.
- **Complexidade:** O(n log n).

### BFS e DFS:
- **Aplicações:** Encontrar caminhos, componentes conectados, verificar ciclos em grafos.
- **Complexidade:** O(V + E), onde V é o número de vértices e E é o número de arestas.

---

## Conclusão
Este roadmap cobre os fundamentos e algoritmos essenciais para dominar DSA em Go. Pratique implementando cada algoritmo e compreendendo suas complexidades. Use o site [Go by Example](https://gobyexample.com) para exemplos práticos e detalhados sobre a linguagem.
