FROM golang:1.19.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# COPY . .

# RUN go build -o main .

# CMD ["/app/main"]

# desenvolvimento com docker-compose exec -it simulador bash. 
# Para rodar o projeto no modo produção, basta comentar a linha abaixo e descomentar as linhas acima.
CMD ["tail", "-f", "/dev/null"] 
