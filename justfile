
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
   
test-cypress-golang-html-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-go-inmemory" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-html-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-html-go-inmemory:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-go-inmemory" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-html-frontend.yaml \
        -f compose.go-backend.yaml \
        -f compose.db-inmemory.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-html-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-go-mongo" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-html-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-html-go-mongo:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-go-mongo" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-html-frontend.yaml \
        -f compose.go-backend.yaml \
        -f compose.db-mongo.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-html-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-js-inmemory" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-html-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-html-js-inmemory:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-js-inmemory" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-html-frontend.yaml \
        -f compose.js-backend.yaml \
        -f compose.db-inmemory.yaml \
        -f compose.expose-ports.yaml \
        up --build
   
test-cypress-golang-html-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-js-mongo" docker compose \
     -f docker-compose.yaml \
     -f compose.golang-html-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e

start-golang-html-js-mongo:
     COMPOSE_PROJECT_NAME="kadai-cypress-golang-html-js-mongo" docker compose \
        -f docker-compose.yaml \
        -f compose.golang-html-frontend.yaml \
        -f compose.js-backend.yaml \
        -f compose.db-mongo.yaml \
        -f compose.expose-ports.yaml \
        up --build
   