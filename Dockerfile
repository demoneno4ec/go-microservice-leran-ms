FROM golang:1.19.0

WORKDIR /application

COPY . .

RUN go mod vendor
RUN go build ./cmd/main.go
