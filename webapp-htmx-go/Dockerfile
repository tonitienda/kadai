# Stage 1: Build the Go binary
FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /app

COPY . .

RUN cd cmd && go build -o ../server

# Stage 2: Create a minimal runtime image
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates

EXPOSE 8080


HEALTHCHECK --interval=10s --timeout=1s --retries=3 CMD wget --tries=1 --spider http://localhost:8080/healthz || exit 1

CMD ["./server"]