FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /gif-to-mp4-convert-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

ENTRYPOINT ["/gif-to-mp4-convert-service/main"]