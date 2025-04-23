# syntax=docker/dockerfile:1

# Stage 1: Build
FROM golang:1.22.5-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

# Stage 2: Run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env

EXPOSE 3000

CMD ["./main"]