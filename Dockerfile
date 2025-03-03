FROM golang:1.23.6-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk --no-cache add ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:latest

WORKDIR /app/server

# Instala dependências para leitura do .env
RUN apk --no-cache add ca-certificates

# Copia o binário da etapa de build
COPY --from=builder /app/server .

# Copia o arquivo .env
# COPY .env .env

EXPOSE 3333

CMD ["sh", "-c", "export $(cat .env | xargs) && ./server"]

# FROM scratch 
# COPY --from=builder /app/server /server
# EXPOSE 3333

# CMD ["/server"]
