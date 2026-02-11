FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init my-http-server && go build -o server .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]