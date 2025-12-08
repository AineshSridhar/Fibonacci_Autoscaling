FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fib-service



FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/fib-service /app/

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/app/fib-service"]