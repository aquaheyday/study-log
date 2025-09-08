# Builder
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o chat-server .

# Runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/chat-server .
EXPOSE 8080
CMD ["./chat-server"]
