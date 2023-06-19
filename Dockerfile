FROM golang:1.20.5-alpine3.18

WORKDIR /app/
COPY csv2json.go /app/

RUN go build csv2json.go;
