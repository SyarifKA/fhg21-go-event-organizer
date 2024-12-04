# FROM golang:1.23-alpine

# WORKDIR /app

# COPY . /app/

# RUN go mod tidy && go build -o binary

# EXPOSE 8888

# ENTRYPOINT [ "/app/binary" ]

# FROM golang:1.23-alpine AS build

# WORKDIR /folderBuild

# COPY . /folderBuild/

# RUN go mod tidy && go build -o /folderBuild/backend

# FROM alpine:3.20 AS prod

# WORKDIR /folderApp

# COPY --from=build /folderBuild /folderApp

# ENV PATH="/folderApp:${PATH}"

# EXPOSE 8080

# ENTRYPOINT [ "backend" ]

FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . /app/

RUN go mod tidy && go build -o /app/backend

FROM alpine:3.20 AS prod

WORKDIR /app

COPY --from=build /app /app

ENV PATH="/app:${PATH}"

EXPOSE 8080

ENTRYPOINT [ "backend" ]