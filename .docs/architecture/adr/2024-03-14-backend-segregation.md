# Separation of HTMX and JSON/Rest endpoints

## Context and Problem Statement

In this project we want to offer different webapplications and clients in general.
Some will request and send JSON data via a REST API, other will be based on HTMX and will request HTML fragments.

For optimization reasons we might want to offer each client a different set of endpoints (for example WebApp or Mobile App). But this analysis and decision is not part of this ADR. We will start with the assumption that the endpoints are the same for now.

## Considered Options

### Use mimetypes to distinguish between JSON and HTML requests

One option is to use the `Accept` header to distinguish between JSON and HTML requests. This is a common approach and is used by many frameworks.

This solution assumes that the endpoints required for the JSON clients and the HTMX clients are the same.

### Use different endpoints for JSON and HTMX requests

Another option is to use different endpoints for JSON and HTMX requests. 

This solution would allow to each client to evolve independently.

## Decision Outcome

We will use different endpoints for JSON and HTMX requests since it is very likely that the requirements for each client will evolve independently.
