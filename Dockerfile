FROM golang:1.22-alpine AS builder

RUN mkdir /app
WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./

# Качаем зависимости (они кэшируются)
RUN go mod download

# Копируем остальное
COPY . .

# Собираем проект
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/api/main.go

# Production stage
FROM alpine:latest AS production

WORKDIR /app

# Копируем скомпилированное приложение и миграции
COPY --from=builder /app/app ./app

# Запуск приложения
CMD ["./app"]