# Use Cases

## User

### Register

The user shall be able to register themselves.

### Authenticate

The user shall be able to authenticate themselves.

## Tasks

### List all tasks

The user shall be able to retrieve their tasks.

That includes all tasks where the user has `view` permission regardless of the owner.

### Create a new task

The user shall be able to create a new task.
When the task is created the user will be the owner of the task.

### Delete a task

The user shall be able to delete a task if they have enough permissions to do it.

### Edit a task

The user shall be able to edit a task if they have enough permissions to do it.

### Share a task

The user shall be able to share a task with other users if they have enough permissions to do it.
When the task is shared, the user can assign different permissions (read only, edit, delete, etc)

### Transfer ownership

The user shall be able to transfer the ownership of a task to another user if they have enough permissions to do it. At that moment the new owner will have all permissions granted.
The old owner will keep the existing permissions except the `share` one that will be transferred to the new owner.

### Revoke permissions

The user with the `share` permission shall be able to revoke the permissions granted to other users.

