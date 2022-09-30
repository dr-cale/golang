# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY api-server-gorilla-mux.go ./

RUN go build -o /api-server-gorilla-mux

EXPOSE 8080

CMD [ "/api-server-gorilla-mux" ]
