const fs = require("node:fs");
const {
  webapps,
  backends,
  e2eRunners,
  systemtestRunners,
  dbs,
} = require("./components.json");

const contents = `
flowchart TD

subgraph webapps
  ${webapps
    .map(
      (w) => `
     ${w.id}["${w.name}
     [${w.language}]
     "]`
    )
    .join("\n")}
end

subgraph backends
  ${backends
    .map(
      (b) => `
     ${b.id}["${b.name}
     [${b.language}]"]`
    )
    .join("\n")}
end

subgraph dbs
    ${dbs
      .map(
        (db) => `
       ${db.id}[("${db.name}")]`
      )
      .join("\n")}
end

subgraph e2e
      ${e2eRunners
        .map(
          (e) => `
         ${e.id}["${e.name}
         [${e.language}]"]`
        )
        .join("\n")}
end

subgraph systemtests
      ${systemtestRunners
        .map(
          (s) => `
         ${s.id}["${s.name}
         [${s.language}]"]`
        )
        .join("\n")}
end

webapps -- uses --> backends
backends -- "stores data in" --> dbs

e2e -- tests --> webapps
systemtests -- tests ---> backends


classDef Tier fill:#08427b,color:#ccc,stroke:#333
classDef Service fill:#1168bd,color:#ccc,stroke:#333
classDef Tests fill:#087b42,color:#333,stroke:#333
classDef TestService fill:#11bd68,color:#333,stroke:#333


class e2e,systemtests Tests
class webapps,backends,dbs Tier
class ${[...webapps, ...backends, ...dbs].map((c) => c.id).join(",")} Service
class ${[...e2eRunners, ...systemtestRunners]
  .map((c) => c.id)
  .join(",")} TestService


`;

fs.writeFileSync("diagram.mmd", contents);
