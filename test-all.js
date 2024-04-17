var exec = require("child_process").exec;

const frontEnds = ["next", "golang-html"];
const backends = ["go", "js"];
const testRunners = ["cypress"];
const dbs = ["inmemory"];

for (frontEnd of frontEnds) {
  for (backend of backends) {
    for (db of dbs) {
      for (runner of testRunners) {
        exec(
          runE2ETests(runner, frontEnd, backend, db),
          function (error, stdout, stderr) {
            console.log("stdout: " + stdout);
            // console.log("stderr: " + stderr);
            // if (error !== null) {
            //   console.log("exec error: " + error);
            // }
          }
        );
      }
    }
  }
}

function runE2ETests(runner, frontEnd, backend, db) {
  command = `COMPOSE_PROJECT_NAME="kadai-${runner}-${frontEnd}-${backend}-${db}" docker compose \\
    -f docker-compose.yaml \\
    -f compose.${frontEnd}-frontend.yaml \\
    -f compose.${backend}-backend.yaml \\
    -f compose.db-${db}.yaml \\
    -f compose.e2e-${runner}.yaml \\
    up --build --exit-code-from e2e`;

  console.log(command);
  exec(command, function (error, stdout, stderr) {
    console.log("stdout: " + stdout);
    // console.log("stderr: " + stderr);
    // if (error !== null) {
    //   console.log("exec error: " + error);
    // }
  });
}
