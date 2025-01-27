# Usa uma imagem base do Go
FROM golang:1.20-alpine

# Cria o diretório de trabalho
WORKDIR /app

# Copia go.mod e go.sum (se existirem)
COPY go.mod  ./


# Copia o restante do código para o container
COPY . .

# Compila a aplicação
RUN go build -o /loadtest main.go

# Define a entrada padrão do container
ENTRYPOINT ["/loadtest"]
