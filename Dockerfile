FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/api/main.go

FROM alpine:latest AS production
WORKDIR /app

COPY --from=builder /app/app ./app

CMD ["./app"]