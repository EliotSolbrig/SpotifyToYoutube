FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

COPY *./ ./

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./main

FROM alpine

RUN apk update


WORKDIR /app
EXPOSE 80
CMD ["./main"]
