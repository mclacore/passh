FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY main.go go.mod ./

RUN go mod download

RUN go build -o passh .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/passh .

RUN chmod +x ./passh

ENTRYPOINT ["./passh"]
