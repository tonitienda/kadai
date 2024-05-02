
unit-test:
  @echo "No unit tests found"

   
start-nextjs-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-nextjs-go-inmemory" docker compose \
      -f compose.frontend-nextjs.yaml \
      -f compose.backend-go.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-nextjs-go-mongo" docker compose \
      -f compose.frontend-nextjs.yaml \
      -f compose.backend-go.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-nextjs-js-inmemory" docker compose \
      -f compose.frontend-nextjs.yaml \
      -f compose.backend-js.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-nextjs-js-mongo" docker compose \
      -f compose.frontend-nextjs.yaml \
      -f compose.backend-js.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-nextjs-clojure-inmemory" docker compose \
      -f compose.frontend-nextjs.yaml \
      -f compose.backend-clojure.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-nextjs-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-nextjs-clojure-mongo" docker compose \
      -f compose.frontend-nextjs.yaml \
      -f compose.backend-clojure.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-htmx-go-go-inmemory" docker compose \
      -f compose.frontend-htmx-go.yaml \
      -f compose.backend-go.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-htmx-go-go-mongo" docker compose \
      -f compose.frontend-htmx-go.yaml \
      -f compose.backend-go.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-htmx-go-js-inmemory" docker compose \
      -f compose.frontend-htmx-go.yaml \
      -f compose.backend-js.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-htmx-go-js-mongo" docker compose \
      -f compose.frontend-htmx-go.yaml \
      -f compose.backend-js.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-htmx-go-clojure-inmemory" docker compose \
      -f compose.frontend-htmx-go.yaml \
      -f compose.backend-clojure.yaml \
      -f compose.db-inmemory.yaml \
      -f compose.expose-ports.yaml \
      up --build
    

   
start-htmx-go-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-htmx-go-clojure-mongo" docker compose \
      -f compose.frontend-htmx-go.yaml \
      -f compose.backend-clojure.yaml \
      -f compose.db-mongo.yaml \
      -f compose.expose-ports.yaml \
      up --build
    


ci-cypress-nextjs-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-go-inmemory" docker compose \
     --profile e2e \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-nextjs-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-go-inmemory" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-nextjs-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-go-mongo" docker compose \
     --profile e2e \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-nextjs-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-go-mongo" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-nextjs-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-js-inmemory" docker compose \
     --profile e2e \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-nextjs-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-js-inmemory" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-nextjs-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-js-mongo" docker compose \
     --profile e2e \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-nextjs-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-js-mongo" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-nextjs-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-clojure-inmemory" docker compose \
     --profile e2e \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-nextjs-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-clojure-inmemory" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-nextjs-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-clojure-mongo" docker compose \
     --profile e2e \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-nextjs-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-nextjs-clojure-mongo" docker compose \
     -f compose.frontend-nextjs.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-htmx-go-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-go-inmemory" docker compose \
     --profile e2e \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-htmx-go-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-go-inmemory" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-htmx-go-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-go-mongo" docker compose \
     --profile e2e \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-htmx-go-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-go-mongo" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-htmx-go-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-js-inmemory" docker compose \
     --profile e2e \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-htmx-go-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-js-inmemory" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-htmx-go-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-js-mongo" docker compose \
     --profile e2e \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-htmx-go-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-js-mongo" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-htmx-go-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-clojure-inmemory" docker compose \
     --profile e2e \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-htmx-go-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-clojure-inmemory" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e



ci-cypress-htmx-go-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-clojure-mongo" docker compose \
     --profile e2e \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     -f compose.ci.yaml \
     up --no-build --exit-code-from e2e

test-cypress-htmx-go-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-cypress-htmx-go-clojure-mongo" docker compose \
     -f compose.frontend-htmx-go.yaml \
     -f compose.backend-clojure.yaml \
     -f compose.db-mongo.yaml \
     -f compose.e2e-cypress.yaml \
     up --build --exit-code-from e2e


ci-bdd-go-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-inmemory" docker compose \
    --profile system \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     -f compose.ci.yaml \
     up  --no-build --exit-code-from system

test-bdd-go-go-inmemory:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-inmemory" docker compose \
     -f compose.backend-go.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from system


ci-bdd-go-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-mongo" docker compose \
    --profile system \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     -f compose.ci.yaml \
     up  --no-build --exit-code-from system

test-bdd-go-go-mongo:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-go-mongo" docker compose \
     -f compose.backend-go.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from system


ci-bdd-go-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-inmemory" docker compose \
    --profile system \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     -f compose.ci.yaml \
     up  --no-build --exit-code-from system

test-bdd-go-js-inmemory:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-inmemory" docker compose \
     -f compose.backend-js.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from system


ci-bdd-go-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-mongo" docker compose \
    --profile system \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     -f compose.ci.yaml \
     up  --no-build --exit-code-from system

test-bdd-go-js-mongo:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-js-mongo" docker compose \
     -f compose.backend-js.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from system


ci-bdd-go-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-clojure-inmemory" docker compose \
    --profile system \
     -f compose.backend-clojure.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     -f compose.ci.yaml \
     up  --no-build --exit-code-from system

test-bdd-go-clojure-inmemory:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-clojure-inmemory" docker compose \
     -f compose.backend-clojure.yaml \
     -f compose.db-inmemory.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from system


ci-bdd-go-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-clojure-mongo" docker compose \
    --profile system \
     -f compose.backend-clojure.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     -f compose.ci.yaml \
     up  --no-build --exit-code-from system

test-bdd-go-clojure-mongo:
  COMPOSE_PROJECT_NAME="kadai-bdd-go-clojure-mongo" docker compose \
     -f compose.backend-clojure.yaml \
     -f compose.db-mongo.yaml \
     -f compose.system-bdd-go.yaml \
     up --build --exit-code-from system

