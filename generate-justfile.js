const fs = require("node:fs");

const frontEnds = ["next", "golang-htmx"];
const backends = ["go", "js"];
const e2eRunners = ["cypress"];
const apiRunners = ["bdd-go"];
const dbs = ["inmemory", "mongo"];

const makeJustE2ETestTasks = ({ frontend, backend, runner, db }) => `
test-${runner}-${frontend}-${backend}-${db}:
  docker compose \\
  -f compose.${frontend}-frontend.yaml \\
  -f compose.${backend}-backend.yaml \\
  -f compose.db-${db}.yaml \\
  -f compose.e2e-${runner}.yaml \\
  config

  COMPOSE_PROJECT_NAME="kadai-${runner}-${frontend}-${backend}-${db}" docker compose \\
     -f compose.${frontend}-frontend.yaml \\
     -f compose.${backend}-backend.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.e2e-${runner}.yaml \\
     up --build --exit-code-from e2e

`;

const makeJustApiTestTasks = ({ backend, runner, db }) => `
test-${runner}-${backend}-${db}:
  docker compose \\
  -f compose.${backend}-backend.yaml \\
  -f compose.db-${db}.yaml \\
  -f compose.e2e-${runner}.yaml \\
  config

  COMPOSE_PROJECT_NAME="kadai-${runner}-${backend}-${db}" docker compose \\
     -f compose.${backend}-backend.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.e2e-${runner}.yaml \\
     up --build --exit-code-from e2e

`;

const makeJustStartTasks = ({ frontend, backend, db }) => `
   
start-${frontend}-${backend}-${db}:
  docker compose \\
  -f compose.${frontend}-frontend.yaml \\
  -f compose.${backend}-backend.yaml \\
  -f compose.db-${db}.yaml \\
  config

  COMPOSE_PROJECT_NAME="kadai-${frontend}-${backend}-${db}" docker compose \\
      -f compose.${frontend}-frontend.yaml \\
      -f compose.${backend}-backend.yaml \\
      -f compose.db-${db}.yaml \\
      -f compose.expose-ports.yaml \\
      up --build
    
`;

const apiCombinations = backends.flatMap((b) =>
  dbs.flatMap((db) => ({
    backend: b,
    db: db,
  }))
);

const appCombinations = frontEnds.flatMap((f) =>
  apiCombinations.map((combination) => ({
    ...combination,
    frontend: f,
  }))
);

const apiTestsCombinations = apiCombinations.flatMap((a) =>
  apiRunners.flatMap((r) => ({
    ...a,
    runner: r,
  }))
);

const e2eTestsCombinations = appCombinations.flatMap((a) =>
  e2eRunners.flatMap((r) => ({
    ...a,
    runner: r,
  }))
);

const contents =
  appCombinations.reduce(
    (contents, combination) => contents + makeJustStartTasks(combination),
    ""
  ) +
  e2eTestsCombinations.reduce(
    (contents, combination) => contents + makeJustE2ETestTasks(combination),
    ""
  ) +
  apiTestsCombinations.reduce(
    (contents, combination) => contents + makeJustApiTestTasks(combination),
    ""
  );

fs.writeFileSync("justfile", contents);
