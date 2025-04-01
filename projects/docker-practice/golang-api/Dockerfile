FROM golang:1.20 AS builder

WORKDIR /app

COPY ./app ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]

EXPOSE 8080
