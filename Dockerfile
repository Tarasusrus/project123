# Установка стадии сборки
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем оставшиеся файлы и собираем проект
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -v -o base-api ./cmd/api

# Финальная стадия
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/base-api /app/base-api
COPY ./configs /app/configs

CMD ["/app/base-api"]