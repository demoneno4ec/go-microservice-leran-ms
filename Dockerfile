FROM golang:1.19.0

WORKDIR /application

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go build ./cmd/main.go

CMD ["air", "-c", ".air.toml"]