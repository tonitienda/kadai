# Response format for REST endpoints

## Context and Problem Statement

When the server returns a resource (for example upon creation) the client will usually need to know the actions available for that resource or where to find the resource details.

One way to achieve this is by sharing the routing and business logic between Client and Server.

For example if customers are retrieved via the `customers` endpoint the client will know that the customer with id `1` can be retrieved via the `customers/1` endpoint.

But also that the `Edition` of the customer can be done using the `PATCH` verb in the `customers/1` endpoint.

Things get more complicated when the client needs to know the actions available for a resource given the status of the resources or the user permissions.

## Considered Options

### Duplicate logic in client and server

One option is to duplicate the routing and business logic in the client and server.
If the Server is stable and follows REST principles if should be straightforward to implement the routing logic in the client and it would follow standard REST practices and not specific server logic.

Still disable/hide links/buttons based on user permissions will require to duplicate the logic.
Non CRUD operations that require specific endpoints and verbs will also need to be made explicit in clients and server.

### HATEOAS

Another option is to use HATEOAS (Hypermedia as the Engine of Application State) to provide the client with the actions available for a resource.

There are different formats to do so like HAL, JSON-LD, Siren, ... but there is no standard format.

None of the formats detail the request that needs to be sent to the server to perform the action. Most examples focus on navigations between resources and not so much about operations and verbs other than GET.

## Decision Outcome

Although HATEOAS offers a way to dedup logic between server and client it is still not clear how to use it or how stable / dynamic it will be, so for now, in v0 of the api, we will return simple JSON responses with the entity's data.

In case of a POST operation (entity creation) we will return the entity's ID as part of the response.
