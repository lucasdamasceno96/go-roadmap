# Dockerizing Golang Applications

Este guia ensina como criar imagens Docker para suas aplicações em Golang. O processo inclui a criação de um arquivo Dockerfile e instruções para construir e executar o contêiner.

## Passo 1: Preparar sua aplicação Golang

Certifique-se de que sua aplicação Go está funcionando corretamente. Por exemplo, considere o seguinte código simples:

```go
// main.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Docker!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}
```

## Passo 2: Criar um Dockerfile

Crie um arquivo chamado `Dockerfile` na mesma pasta onde está o arquivo `main.go`. O Dockerfile define como sua aplicação será empacotada.

### Exemplo de Dockerfile

```dockerfile
# Usar uma imagem oficial do Golang como base
FROM golang:1.20 AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos da aplicação para o contêiner
COPY . .

# Baixar dependências e compilar o binário
RUN go mod tidy \
    && go build -o main .

# Usar uma imagem mínima para o ambiente de produção
FROM debian:bullseye-slim

# Copiar o binário da fase de build para a fase final
COPY --from=builder /app/main /app/main

# Definir o diretório de trabalho
WORKDIR /app

# Expor a porta que será usada pela aplicação
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["/app/main"]
```

## Passo 3: Construir a imagem Docker

Com o Docker instalado em sua máquina, abra um terminal na pasta do projeto e execute o comando:

```bash
docker build -t my-golang-app .
```

Neste exemplo:
- `-t my-golang-app`: Define o nome e a tag da imagem Docker.
- `.`: Especifica o diretório atual como contexto de build.

## Passo 4: Executar o contêiner

Para executar o contêiner com a imagem criada, use o comando:

```bash
docker run -d -p 8080:8080 --name golang-container my-golang-app
```

Aqui:
- `-d`: Executa o contêiner em segundo plano.
- `-p 8080:8080`: Mapeia a porta 8080 do contêiner para a porta 8080 do host.
- `--name golang-container`: Define um nome para o contêiner.
- `my-golang-app`: Nome da imagem que será usada.

Acesse sua aplicação no navegador ou com `curl` em `http://localhost:8080`.

## Passo 5: Publicar a imagem (opcional)

Se quiser compartilhar a imagem, você pode enviá-la para o Docker Hub:

1. Faça login no Docker Hub:
   ```bash
   docker login
   ```

2. Marque a imagem com seu nome de usuário:
   ```bash
   docker tag my-golang-app username/my-golang-app
   ```

3. Envie a imagem para o Docker Hub:
   ```bash
   docker push username/my-golang-app
   ```

Agora a imagem está disponível para outros usuários baixarem e utilizarem.

---

## Conclusão

Parabéns! Você aprendeu a criar imagens Docker para suas aplicações Golang. Use esse processo para simplificar o deployment e manter consistência entre ambientes. Se precisar de mais ajuda, não hesite em perguntar!
