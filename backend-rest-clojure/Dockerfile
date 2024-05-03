FROM clojure:latest as build

# Set the working directory in the container
WORKDIR /app



# TODO: Use a lightweight base image
FROM clojure:latest

WORKDIR /app

COPY . .

ENTRYPOINT  ["lein", "ring", "serve-headless", "8080"]