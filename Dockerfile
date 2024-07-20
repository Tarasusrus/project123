FROM golang:1.22-alpine AS builder

RUN mkdir /app
ADD .. /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/api/main.go

FROM alpine:latest AS production
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/app ./app

CMD ["./app"]
