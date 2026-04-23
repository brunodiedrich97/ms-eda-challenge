FROM golang:1.26.2

WORKDIR /app/

# Instala a dependência do Kafka
RUN apt-get update && apt-get install -y librdkafka-dev

# Copia os arquivos de dependência primeiro (otimiza o cache do Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Compila a aplicação
RUN go build -o main ./cmd/walletcore/main.go

# Expõe a porta (8080 para wallet, 3003 para balance)
EXPOSE 8080

# Executa a aplicação
CMD ["./main"]

# container ficar de pé
# CMD ["tail", "-f", "/dev/null"]