# Используем официальный образ Go
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o runner ./cmd/bot/

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/runner .

COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./runner"]