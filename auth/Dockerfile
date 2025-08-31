FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

COPY auth/go.mod auth/go.sum ./auth/
COPY common/go.mod common/go.sum ./common/

#
WORKDIR /app/auth

RUN go mod download

# 
WORKDIR /app

COPY auth auth
COPY common common

WORKDIR /app/auth
ENV CGO_ENABLED=false
ENV GOOSE=linux
RUN go build -o auth cmd/auth/main.go

FROM alpine:3.22
WORKDIR /app

COPY --from=builder /app/auth/auth .
COPY --from=builder /app/auth/migrations ./migrations
COPY swagger swagger

CMD ["./auth"]
