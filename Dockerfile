FROM golang:1.20.12-alpine3.19

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go run .