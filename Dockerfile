FROM golang:1.21.5-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN apk add --no-cache \
    ca-certificates \
    git

COPY . .

RUN go mod download

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .


  
COPY prueba.txt /app/prueba.txt.


EXPOSE 8080

CMD ["./main"]
