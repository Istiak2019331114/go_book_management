FROM golang:1.24 AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /go-book ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go-book .

ENTRYPOINT ["./go-book", "serve"]
