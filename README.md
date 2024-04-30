# kadai

Application to manage tasks used to practice, test, update and showcase my work.

Find the documentation under the `.docs` folder.

Quick links:

- [Requirements](.docs/requirements/use-cases.md)
- [Architecture](.docs/architecture/software-architecture-document.md)
  - [Decision Records](.docs/architecture/adr)
- [Backlog](.docs/tasks/backlog.md)

## Tech Stack

In real world projects I would never define the list of tech stack before understanding the requirements and the constraints of the projects. 
I would also try to keep the stack simple, short and [boring](https://mcfunley.com/choose-boring-technology). See my own [blog post](https://tsoobame.github.io/blog/resume-driven-development.html) about it.

However, this project is a playground to test and learn new technologies so I can risk it and any technology I feel like.
Here I am listing some of the technologies I am planning to use in this project and some of the patterns I want to try out.
This is a sort of wishlist. The list will keep growing and I will not be able to use/learn everything listed there.

I am marking with ✔️ the technologies/approaches I am using and with 🚧 the ones I started to add but are still not functional.

### Technologies

#### Backend

- Golang ✔️
- Java with Spring Boot
- Kotlin
- C#
- Nodejs with Express ✔️
- Nodejs with Nestjs
- Clojure
- Python / FastAPI
- Python / Django

#### Frontend 

- React with Nextjs ✔️
- React with Remix 🚧
- HTMX with Golang ✔️
- HTMX with Nodejs/Ejs

#### APIs

- Graphql
- RESTful API ✔️
- gRPC
- Websockets
- Server Sent Events

#### Acceptance / E2E Testing

- BDD
- Cypress ✔️ (AFAIK not suitable for tests that involve multiple browsers / users) 
- Test Cafe
- Selenium

#### Message Brokers

- Kafka
- RabbitMQ
- NATS

#### Databases

- Postgresql
- Mysql
- MongoDB 🚧
- Neo4j
- Datomic
- Sqlite ?
- Cassandra ?
- Snowflake ?
- Databricks ?


#### Cache

- Redis
- Memcached
- Hazelcast

#### Containers/Orchestration

- Docker ✔️
- Podman
- Docker Compose ✔️ / Swarm
- Kubernetes
- Kustomize
- Telepresence
- Helm

### Infras as Code

- Terraform
- Ansible
- Chef
- Puppet

### WebServers

- Nginx
- Apache

### CI/CD

- Jenkins
- Github Actions ✔️
- Gitlab CI/CD
- ArgoCD
- FluxCD
- CircleCI


#### Monitoring

- OpenTelemetry
- OpenTracing
- OpenMetrics
- OpenCensus
- OpenAPM
- Prometheus
- Grafana
- Jaeger
- Zipkin
- ELK
- Datadog
- Nagios

### Cloud/Computing

- AWS
- GCP
- Azure
- Hadoop
- DigitalOcean

### UX Patterns

- Command -> Undo: 🚧 
  for example instead of asking for confirmation when deleting a task I will offer the option to undo the change (similar to Gmail).


### Geospatial

- Tile38 (not sure how it will fit the TODO application yet)
