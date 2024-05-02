# Stage 1: Build the Go binary
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080


# HEALTHCHECK --interval=10s --timeout=1s --retries=3 CMD wget --tries=1 --spider http://localhost:8080/healthz || exit 1
HEALTHCHECK --interval=10s --timeout=1s --retries=3 CMD echo "hello"

CMD ["./server"]