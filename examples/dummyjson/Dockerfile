FROM golang:1.23-bookworm

RUN go install github.com/mitranim/gow@latest

WORKDIR examples/dummyjson

COPY . .

EXPOSE 5000

ENTRYPOINT ["gow", "run", "main.go"]
