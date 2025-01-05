FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY *./ ./

RUN CGO_ENABLED=0 go build -o ./main

FROM alpine

RUN apk update


WORKDIR /app
EXPOSE 8080
CMD ["./main"]
