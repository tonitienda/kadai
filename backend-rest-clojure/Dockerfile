# TODO: Use a lightweight base image
FROM clojure:latest

WORKDIR /app

RUN apt update && apt upgrade && \
    apt install wget

COPY . .

HEALTHCHECK --interval=10s --timeout=1s --retries=3 CMD wget --tries=1 --spider http://localhost:8080/healthz || exit 1

CMD  ["clj", "-X", "kadai.core/start-server"]