# Use a imagem oficial do Golang para buildar o projeto
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copia o main.go para o container
COPY main.go .

# Compila o binário
RUN go build -o stress-test main.go

# Imagem final, mais enxuta
FROM alpine:latest

WORKDIR /app

# Copia o binário compilado
COPY --from=builder /app/stress-test .

# Expõe a porta 8080 (ajuste se necessário)
EXPOSE 8080

# Comando de entrada, aceita parâmetros via linha de comando
ENTRYPOINT ["./stress-test"]