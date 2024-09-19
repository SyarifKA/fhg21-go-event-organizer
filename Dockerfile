# FROM golang:1.23

# WORKDIR /app

# COPY . /app/

# RUN go mod tidy

# ENTRYPOINT go run main.go

# CMD [ "go", "run", "main.go" ]

FROM golang:1.23-alpine

WORKDIR /app

COPY . /app/

RUN go mod tidy && go build -o binary

EXPOSE 8888

ENTRYPOINT [ "/app/binary" ]