FROM golang:latest AS builder

RUN apt-get update 

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOPATH/backend

COPY . .

COPY go.mod .

RUN go mod download

RUN ls


RUN go build src/main.go

ENV PYTHONUNBUFFERED 1

CMD ["./main"]