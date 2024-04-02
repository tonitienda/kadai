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

### Technologies

#### Backend

- Golang
- Java with Spring Boot
- Kotlin
- C#
- Nodejs with Express
- Nodejs with Nestjs
- Clojure

#### Frontend 

- React with Nextjs
- React with Remix
- HTMX with Golang
- HTMX with Nodejs/Ejs

#### APIs

- Graphql
- RESTful API
- gRPC
- Websockets
- Server Sent Events

#### Message Brokers

- Kafka
- RabbitMQ
- NATS

#### Databases

- Postgresql
- Mysql
- MongoDB


#### Cache

- Redis
- Memcached
- Hazelcast

#### DevOps

- Docker
- Podman
- Docker swarm
- Kubernetes
- Kustomize
- Telepresence
- Helm
- Terraform
- Ansible
- Jenkins
- Github Actions
- Gitlab CI/CD
- ArgoCD
- FluxCD


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

### UX Patterns

- Command -> Undo: for example instead of asking for confirmation when deleting a task I will offer the option to undo the change (similar to Gmail).

