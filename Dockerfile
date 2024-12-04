FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o go-api /app/src/api/main.go

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/go-api .
EXPOSE 8081
CMD ["./go-api"]