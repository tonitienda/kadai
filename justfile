
unit-test:
  @echo "No unit tests found"

   
start-nextjs-go-inmemory:
  docker compose \
  -f compose.nextjs-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-nextjs-go-inmemory" docker compose \
      -f compose.nextjs-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-go-mongo:
  docker compose \
  -f compose.nextjs-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-nextjs-go-mongo" docker compose \
      -f compose.nextjs-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-js-inmemory:
  docker compose \
  -f compose.nextjs-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-nextjs-js-inmemory" docker compose \
      -f compose.nextjs-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-js-mongo:
  docker compose \
  -f compose.nextjs-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-nextjs-js-mongo" docker compose \
      -f compose.nextjs-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-go-inmemory:
  docker compose \
  -f compose.htmx-go-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-htmx-go-go-inmemory" docker compose \
      -f compose.htmx-go-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-go-mongo:
  docker compose \
  -f compose.htmx-go-frontend.yaml \
  -f compose.go-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-htmx-go-go-mongo" docker compose \
      -f compose.htmx-go-frontend.yaml \
      -f compose.go-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-js-inmemory:
  docker compose \
  -f compose.htmx-go-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-inmemory.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-htmx-go-js-inmemory" docker compose \
      -f compose.htmx-go-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-js-mongo:
  docker compose \
  -f compose.htmx-go-frontend.yaml \
  -f compose.js-backend.yaml \
  -f compose.db-mongo.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-htmx-go-js-mongo" docker compose \
      -f compose.htmx-go-frontend.yaml \
      -f compose.js-backend.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

test-cypress-nextjs-go-inmemory:
  docker compose \
  -f compose.frontend-nextjs.yaml \
  -f compose.backend-go.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-go-inmemory" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-nextjs-go-mongo:
  docker compose \
  -f compose.frontend-nextjs.yaml \
  -f compose.backend-go.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-go-mongo" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-nextjs-js-inmemory:
  docker compose \
  -f compose.frontend-nextjs.yaml \
  -f compose.backend-js.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-js-inmemory" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-nextjs-js-mongo:
  docker compose \
  -f compose.frontend-nextjs.yaml \
  -f compose.backend-js.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-js-mongo" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-htmx-go-go-inmemory:
  docker compose \
  -f compose.frontend-htmx-go.yaml \
  -f compose.backend-go.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-go-inmemory" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-htmx-go-go-mongo:
  docker compose \
  -f compose.frontend-htmx-go.yaml \
  -f compose.backend-go.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-go-mongo" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-htmx-go-js-inmemory:
  docker compose \
  -f compose.frontend-htmx-go.yaml \
  -f compose.backend-js.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-js-inmemory" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-cypress-htmx-go-js-mongo:
  docker compose \
  -f compose.frontend-htmx-go.yaml \
  -f compose.backend-js.yaml \
  -f compose.db-mongo.yaml \
  -f compose.e2e-cypress.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-js-mongo" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


test-bdd-go-go-inmemory:
  docker compose \
  -f compose.backend-go.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.system-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-inmemory" docker compose \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from e2e


test-bdd-go-go-mongo:
  docker compose \
  -f compose.backend-go.yaml \
  -f compose.db-mongo.yaml \
  -f compose.system-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-mongo" docker compose \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from e2e


test-bdd-go-js-inmemory:
  docker compose \
  -f compose.backend-js.yaml \
  -f compose.db-inmemory.yaml \
  -f compose.system-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-inmemory" docker compose \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from e2e


test-bdd-go-js-mongo:
  docker compose \
  -f compose.backend-js.yaml \
  -f compose.db-mongo.yaml \
  -f compose.system-bdd-go.yaml \
  config

  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-mongo" docker compose \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from e2e

