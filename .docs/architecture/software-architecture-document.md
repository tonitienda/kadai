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

Auth0["Auth0
[Software System]

Software that handles the authentication and authorization of the users."]

User -- "manages tasks using" --> Kadai
Kadai -- "handles authentication using" --> Auth0


classDef person fill:#08427b
classDef system fill:#1168bd
classDef external fill:#999999

class User person
class Kadai system
class Auth0 external
```

### 2.2 Kadai System Architecture

```mermaid
flowchart TB

User["Normal User
[Person]"]

subgraph Kadai

ReactWebApp["WebApp
[Software System]

Web application to manage tasks.
[React]"]

HtmxWebApp["WebApp
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

Auth0["Auth0
[Software System]

Software that handles the authentication and authorization of the users."]

User -- "manages tasks using" --> ReactWebApp
User -- "manages tasks using" --> HtmxWebApp
User -- "manages tasks using" --> MobileApp

ReactWebApp -- "communicates with" --> Backend
HtmxWebApp -- "communicates with" --> Backend
MobileApp -- "communicates with" --> Backend

ReactWebApp -- "gets user token using" --> Auth0
HtmxWebApp -- "gets user token using" --> Auth0
MobileApp -- "gets user token using" --> Auth0


classDef person fill:#08427b
classDef system fill:#1168bd
classDef external fill:#999999

class User person
class ReactWebApp,HtmxWebApp,MobileApp,Backend system
class Auth0 external
```

### 2.3 Backend Architecture

The backend is a monolithic application that is divided into the following layers:

- **HTTP Layer**: This layer is responsible for handling the HTTP requests and responses. It is the entry point of the application.
- **Core**: This layer is responsible for handling the business logic of the application.
- **Data Access**: This layer is responsible for handling the communication with the database.

The backend is also split into vertical slices, where each slice is responsible for handling a specific feature of the application.

For now only the Tasks feature is implemented.

```mermaid
flowchart TB

ReactWebApp["WebApp
[Software System]

Web application to manage tasks.
[React]"]

HtmxWebApp["WebApp
[Software System]

Web application to manage tasks.
[HTMX]"]

MobileApp["MobileApp
[Software System]

Mobile application to manage tasks.
[React Native]"]


subgraph Backend
    AuthMw["Auth Middleware
    [Software System]

    Handles the authentication of the users.
    [Golang]"]

    JsonApi["JSON API
    [Software System]

    Handles the HTTP requests and responses in JSON format.
    [Golang / JSON]"]

    HTMXApi["HTMX API
    [Software System]

    Handles the HTTP requests and responses in HTMX format.
    [Golang / JSON]"]

    Core["Core
    [Software System]

    Handles the business logic of the application.
    [Golang]"]

    DataAccess["Data Access
    [Software System]

    Handles the communication with the database.
    [Golang / SQL]"]

    DataBase[("Database
    [Software System]

    Stores the data of the application."
    )]


end

Auth0["Auth0
[Software System]

Software that handles the authentication and authorization of the users."]

ReactWebApp -- "communicates with" --> JsonApi
HtmxWebApp -- "communicates with" --> HTMXApi
MobileApp -- "communicates with" --> JsonApi

ReactWebApp -- "gets user token using" --> Auth0
HtmxWebApp -- "gets user token using" --> Auth0
MobileApp -- "gets user token using" --> Auth0


JsonApi -- "manages tasks data using" --> Core
HTMXApi -- "manages tasks data using" --> Core
Core -- "stores and retrieves data using" --> DataAccess
DataAccess -- "connects to" --> DataBase

classDef system fill:#1168bd
classDef external fill:#999999

class AuthMw,JsonApi,HTMXApi,Core,DataAccess,DataBase system
class ReactWebApp,HtmxWebApp,MobileApp,Auth0 external
```

## 2.2 Scenarios

### 2.2.1 Adding and Deleting a Task

#### Description

- The user adds a new task to the system specifying a title and a description via the GUI.
- The GUI sends a POST request to the backend with the specified data.
  - If the request is successful, the backend responds with a 201 - Created status code.
- The user deletes a task via the GUI.
- The GUI sends a DELETE request to the backend with the task id without asking for confirmation.
  - If the request is successful, the backend responds with a 202 - Accepted status code and a JSON object with the action that can be executed to undo the deletion.
  - The GUI offers the option to undo the deletion.
- The user decides to undo the deletion via the GUI.
- The GUI executes the action provided by the server that will undo the deletion.


Ommited the permissions checks for brevity. Users can only undo the actions they executed.

#### View

```mermaid
sequenceDiagram

Actor User
Participant GUI
Participant Backend

User ->> GUI: add_task(title, description)
GUI ->> Backend: POST /tasks {title, description}
Backend ->> GUI: 201 Created

User ->> GUI: delete_task(task_id)
GUI ->> Backend: DELETE /tasks/task_id
Backend ->> GUI: 202 Accepted { _actions: { undo: action }}

alt User wants to undo the deletion
    User ->> GUI: execute(action)
    GUI ->> Backend: POST /actions {action}
    Backend ->> GUI: 202 Accepted
end

```