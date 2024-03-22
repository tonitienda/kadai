# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN cd cmd && go build -o ../server

# Stage 2: Create a minimal runtime image
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]