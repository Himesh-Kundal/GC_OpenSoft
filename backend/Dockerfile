FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates curl

ENV REDIS_ADDR=redis:6379

COPY --from=builder /app/main .
COPY employee.json ./employee.json

EXPOSE 8080

CMD ["./main"]
