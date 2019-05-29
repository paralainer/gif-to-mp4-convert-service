FROM golang:1.12.5-alpine3.9

ENV GO111MODULE=on

RUN apk upgrade -U \
 && apk add ca-certificates ffmpeg libva-intel-driver \
 && apk add --no-cache bash git openssh \
 && rm -rf /var/cache/*

WORKDIR /gif-to-mp4-convert-service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

ENTRYPOINT ["/gif-to-mp4-convert-service/main"]