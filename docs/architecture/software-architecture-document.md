# Software Architecture Document

## 1. Introduction

### 1.1 Purpose

The purpose of this document is to provide a comprehensive overview of the software architecture of the system. It will describe the system's architecture, its components, and the relationships between them.

## 2. Architectural Representation

### 2.1 High-Level Architecture

The system is designed following the client/server architecture. 

The server is a single monolithic application that serves the client requests. The server is responsible for handling the business logic, data storage, and the communication with the clients.

There are two types of clients: the web application and the mobile application. Both clients communicate with the server using HTTP requests.

As for the web applications, and for learning purposes, we will have two different implementations: one using React and another using HTMX. Both implementations will be served by the same server.

```mermaid
flowchart TB

User["Normal User
[Person]

Person that uses Kadai to manage their tasks."]


Kadai["Kadai
[Software System]

Software that allows users to manage their tasks"]

Authentik["Authentik
[Software System]

Software that handles the authentication and authorization of the users."]

User -- "manages tasks using" --> Kadai
Kadai -- "handles authentication using" --> Authentik


classDef person fill:#08427b
classDef system fill:#1168bd
classDef external fill:#999999

class User person
class Kadai system
class Authentik external
```

### 2.2 Kadai System Architecture

```mermaid
flowchart TB

User["Normal User
[Person]"]

subgraph Kadai

WebApp1["WebApp
[Software System]

Web application to manage tasks.
[React]"]

WebApp2["WebApp
[Software System]

Web application to manage tasks.
[HTMX]"]

MobileApp["MobileApp
[Software System]

Mobile application to manage tasks.
[React Native]"]


Backend["Backend
[Software System]

Server that handles the business logic and data storage.
[Golang]"]
end

Authentik["Authentik
[Software System]

Software that handles the authentication and authorization of the users."]

User -- "manages tasks using" --> WebApp1
User -- "manages tasks using" --> WebApp2
User -- "manages tasks using" --> MobileApp

WebApp1 -- "communicates with" --> Backend
WebApp2 -- "communicates with" --> Backend
MobileApp -- "communicates with" --> Backend

Backend -- "handles authentication using" --> Authentik
WebApp1 -- "gets user token using" --> Authentik
WebApp2 -- "gets user token using" --> Authentik
MobileApp -- "gets user token using" --> Authentik


classDef person fill:#08427b
classDef system fill:#1168bd
classDef external fill:#999999

class User person
class WebApp1,WebApp2,MobileApp,Backend system
class Authentik external
```