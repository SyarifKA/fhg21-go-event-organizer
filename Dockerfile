FROM golang:1.23

WORKDIR /app

COPY . /app/

RUN go mod tidy

ENTRYPOINT go run main.go