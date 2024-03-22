# Authorization

## Context and Problem Statement

In Kadai users can manage their own tasks but also can manage tasks from other users if
the owners granted them permission.

Examples of actions that can be performed over tasks are:

view
edit
delete
share
assign
etc.

Only the owner of a task can grant permissions to other users.
The owner can transfer the ownership of a task to another user. The previous owner will keep all the existing permissions except the `share` one that will be transferred to the new owner.

See requirements for more information.

## Considered Options

### Role-based access control

We can use role-based access control to manage permissions. 
We could assign different roles to a pair user/task.

Some examples of roles could be:
- owner
- viewer
- editor

This approach is more static than what we might need.

### Relationship-based access control

In this approach, we can manage permissions based on the relationship between the user and the task. That offers more flexibility than the role-based access control and the granularity we need.

## Decision Outcome

We will use relationship-based access control to manage permissions in Kadai.
The permission control will be always based on the Relation user/task + allowed action.
When a user creates a task, we will keep track of the ownership, but also create the permissions granted to the user. So the access control will only be based on the granted actions.

## Out of scope 

If we need a role of `admin` that has access over all tasks or tasks in an organization or similar, we will address it in a separate design.