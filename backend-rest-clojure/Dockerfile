FROM clojure:latest as build

# Set the working directory in the container
WORKDIR /app



# TODO: Use a lightweight base image
FROM clojure:latest as runtime

WORKDIR /app

COPY . .

ENTRYPOINT  ["lein", "run"]