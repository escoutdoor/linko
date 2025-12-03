FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

COPY driver/go.mod driver/go.sum ./driver/
COPY common/go.mod common/go.sum ./common/

WORKDIR /app/driver

RUN go mod download

WORKDIR /app

COPY driver driver
COPY common common

WORKDIR /app/driver
ENV CGO_ENABLED=false
ENV GOOSE=linux
RUN go build -o driver cmd/driver/main.go

FROM alpine:3.22
WORKDIR /app

COPY --from=builder /app/driver/driver .
COPY --from=builder /app/driver/migrations ./migrations
COPY swagger swagger

CMD ["./driver"]
