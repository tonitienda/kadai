
   
start-next-go-inmemory:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-next-go-inmemory" docker compose \
      -f compose.next-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-next-go-mongo:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-next-go-mongo" docker compose \
      -f compose.next-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-next-js-inmemory:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-next-js-inmemory" docker compose \
      -f compose.next-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-next-js-mongo:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-next-js-mongo" docker compose \
      -f compose.next-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-golang-htmx-go-inmemory:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-golang-htmx-go-inmemory" docker compose \
      -f compose.golang-htmx-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-golang-htmx-go-mongo:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-golang-htmx-go-mongo" docker compose \
      -f compose.golang-htmx-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-golang-htmx-js-inmemory:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-golang-htmx-js-inmemory" docker compose \
      -f compose.golang-htmx-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-golang-htmx-js-mongo:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-golang-htmx-js-mongo" docker compose \
      -f compose.golang-htmx-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

test-cypress-next-go-inmemory:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-next-go-inmemory" docker compose \
     -f compose.next-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-next-go-mongo:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-next-go-mongo" docker compose \
     -f compose.next-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-next-js-inmemory:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-next-js-inmemory" docker compose \
     -f compose.next-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-next-js-mongo:
  docker compose \
  -f compose.next-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-next-js-mongo" docker compose \
     -f compose.next-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-golang-htmx-go-inmemory:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-go-inmemory" docker compose \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-golang-htmx-go-mongo:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-go-mongo" docker compose \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.go-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-golang-htmx-js-inmemory:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-js-inmemory" docker compose \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-golang-htmx-js-mongo:
  docker compose \
  -f compose.golang-htmx-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-golang-htmx-js-mongo" docker compose \
     -f compose.golang-htmx-frontend.yaml \
     -f compose.js-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-bdd-go-go-inmemory:
  docker compose \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-inmemory" docker compose \
     -f compose.go-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-bdd-go.yaml \
     up --build --exit-code-from e2e


test-bdd-go-go-mongo:
  docker compose \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-mongo" docker compose \
     -f compose.go-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-bdd-go.yaml \
     up --build --exit-code-from e2e


test-bdd-go-js-inmemory:
  docker compose \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-inmemory" docker compose \
     -f compose.js-backend.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-bdd-go.yaml \
     up --build --exit-code-from e2e


test-bdd-go-js-mongo:
  docker compose \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-mongo" docker compose \
     -f compose.js-backend.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-bdd-go.yaml \
     up --build --exit-code-from e2e

