# Separation of HTMX/JSON/Rest and Golang/C#/Nodejs endpoints

## Context and Problem Statement

In this project we want to test different techcologies.
We plan to implement different backends in different languages (Golang, C#, Kotlin, Nodejs) that expose Rest/JSON Apis and HTMLX.

## Considered Options

### Split backend by Language

One option is to create one backend for each language and expose the endpoints for JSON and HTMX requests within the same backend.

### Split backend by Technology / Language

Another option is to create one backend for each language and endpoint, creating one backend for each combination of language and technology.

## Decision Outcome

Although this is a lab project, in the real world we would choose one technology and language for each service so we will follow this approach.

We will name each service with the naming convention `backend-<language>-<technology>`.
For example: `backend-rest-go`, `backend-golang-htmx`, `backend-csharp-rest`, ...