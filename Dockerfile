FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./

CMD ["go run main.go"]
