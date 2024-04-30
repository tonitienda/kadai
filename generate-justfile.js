const fs = require("node:fs");

const frontEnds = ["nextjs", "htmx-go"];
const backends = ["go", "js"];
const e2eRunners = ["cypress"];
const systemtests = ["bdd-go"];
const dbs = ["inmemory", "mongo"];

const makeJustE2ETestTasks = ({ frontend, backend, runner, db }) => `

ci-${runner}-${frontend}-${backend}-${db}:
  COMPOSE_PROJECT_NAME="kadai-${runner}-${frontend}-${backend}-${db}" docker compose \\
     --profile e2e \\
     -f compose.frontend-${frontend}.yaml \\
     -f compose.backend-${backend}.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.e2e-${runner}.yaml \\
     -f compose.ci.yaml
     up --build --exit-code-from e2e

test-${runner}-${frontend}-${backend}-${db}:
  COMPOSE_PROJECT_NAME="kadai-${runner}-${frontend}-${backend}-${db}" docker compose \\
     -f compose.frontend-${frontend}.yaml \\
     -f compose.backend-${backend}.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.e2e-${runner}.yaml \\
     up --build --exit-code-from e2e

`;

const makeJustSystemTestTasks = ({ backend, runner, db }) => `
ci-${runner}-${backend}-${db}:
  COMPOSE_PROJECT_NAME="kadai-${runner}-${backend}-${db}" docker compose \\
    --profile system \\
     -f compose.backend-${backend}.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.system-${runner}.yaml \\
     -f compose.ci.yaml \\
     up --build --exit-code-from system

test-${runner}-${backend}-${db}:
  COMPOSE_PROJECT_NAME="kadai-${runner}-${backend}-${db}" docker compose \\
     -f compose.backend-${backend}.yaml \\
     -f compose.db-${db}.yaml \\
     -f compose.system-${runner}.yaml \\
     up --build --exit-code-from system

`;

const makeJustStartTasks = ({ frontend, backend, db }) => `
   
start-${frontend}-${backend}-${db}:
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

const systemTestsCombinations = apiCombinations.flatMap((a) =>
  systemtests.flatMap((r) => ({
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
  `
unit-test:
  @echo "No unit tests found"
` +
  appCombinations.reduce(
    (contents, combination) => contents + makeJustStartTasks(combination),
    ""
  ) +
  e2eTestsCombinations.reduce(
    (contents, combination) => contents + makeJustE2ETestTasks(combination),
    ""
  ) +
  systemTestsCombinations.reduce(
    (contents, combination) => contents + makeJustSystemTestTasks(combination),
    ""
  );

fs.writeFileSync("justfile", contents);
