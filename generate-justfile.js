const fs = require("node:fs");

const frontEnds = ["next", "golang-htmx"];
const backends = ["go", "js"];
const testRunners = ["cypress"];
const dbs = ["inmemory", "mongo"];

const makeJustTasks = ({ frontend, backend, runner, db }) => `
test-${runner}-${frontend}-${backend}-${db}:
  COMPOSE_PROJECT_NAME="kadai-${runner}-${frontend}-${backend}-${db}" docker compose \\
     -f docker-compose.yaml \\
     -f compose.${frontend}-frontend.yaml \\
     -f compose.${backend}-backend.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.e2e-${runner}.yaml \\
     up --build --exit-code-from e2e

start-${frontend}-${backend}-${db}:
     COMPOSE_PROJECT_NAME="kadai-${runner}-${frontend}-${backend}-${db}" docker compose \\
        -f docker-compose.yaml \\
        -f compose.${frontend}-frontend.yaml \\
        -f compose.${backend}-backend.yaml \\
        -f compose.db-${db}.yaml \\
        -f compose.expose-ports.yaml \\
        up --build
   `;

const combinations = frontEnds.flatMap((f) =>
  backends.flatMap((b) =>
    testRunners.flatMap((r) =>
      dbs.flatMap((db) => ({
        frontend: f,
        backend: b,
        runner: r,
        db: db,
      }))
    )
  )
);

console.log(combinations);

const contents = combinations.reduce(
  (contents, combination) => contents + makeJustTasks(combination),
  ""
);

fs.writeFileSync("justfile", contents);
