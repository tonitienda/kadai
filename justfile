
test-cypress-next-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-next-go-inmemory" docker compose \
     -f docker-compose.yaml \
     -f compose.next-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-next-go-inmemory:
     COMPOSE_PROJECT_NAME="kadai-cypress-next-go-inmemory" docker compose \
        -f docker-compose.yaml \
        -f compose.next-frontend.yaml \
        -f compose.go-backend.yaml \
        -f compose.db-inmemory.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-next-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-next-go-mongo" docker compose \
     -f docker-compose.yaml \
     -f compose.next-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-next-go-mongo:
     COMPOSE_PROJECT_NAME="kadai-cypress-next-go-mongo" docker compose \
        -f docker-compose.yaml \
        -f compose.next-frontend.yaml \
        -f compose.go-backend.yaml \
        -f compose.db-mongo.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-next-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-next-js-inmemory" docker compose \
     -f docker-compose.yaml \
     -f compose.next-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-next-js-inmemory:
     COMPOSE_PROJECT_NAME="kadai-cypress-next-js-inmemory" docker compose \
        -f docker-compose.yaml \
        -f compose.next-frontend.yaml \
        -f compose.js-backend.yaml \
        -f compose.db-inmemory.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-next-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-next-js-mongo" docker compose \
     -f docker-compose.yaml \
     -f compose.next-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-next-js-mongo:
     COMPOSE_PROJECT_NAME="kadai-cypress-next-js-mongo" docker compose \
        -f docker-compose.yaml \
        -f compose.next-frontend.yaml \
        -f compose.js-backend.yaml \
        -f compose.db-mongo.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-htmx-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-go-inmemory" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-htmx-go-inmemory:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-go-inmemory" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-htmx-frontend.yaml \
        -f compose.go-backend.yaml \
        -f compose.db-inmemory.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-htmx-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-go-mongo" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-htmx-go-mongo:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-go-mongo" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-htmx-frontend.yaml \
        -f compose.go-backend.yaml \
        -f compose.db-mongo.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-htmx-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-js-inmemory" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-htmx-js-inmemory:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-js-inmemory" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-htmx-frontend.yaml \
        -f compose.js-backend.yaml \
        -f compose.db-inmemory.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-htmx-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-js-mongo" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-htmx-js-mongo:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-js-mongo" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-htmx-frontend.yaml \
        -f compose.js-backend.yaml \
        -f compose.db-mongo.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-all: 
      just test-cypress-next-go-inmemory
      just test-cypress-next-go-mongo
      just test-cypress-next-js-inmemory
      just test-cypress-next-js-mongo
      just test-cypress-golang-htmx-go-inmemory
      just test-cypress-golang-htmx-go-mongo
      just test-cypress-golang-htmx-js-inmemory
      just test-cypress-golang-htmx-js-mongo