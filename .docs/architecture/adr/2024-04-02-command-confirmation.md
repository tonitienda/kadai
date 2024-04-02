# Command confirmation

## Context and Problem Statement

The user can perform actions that have side effects, such as deleting a task (or any other resource). We want to make sure that the user really intended to perform that action and did not execute it by mistake.

## Considered Options

### Confirmation before the action

One way of understaing user intent to perform an action is to ask for confirmation before the action is executed. It can be done, for example, with a modal dialog.

One drawback is that the confirmation dialog is only performed in a GUI and not in API calls.
The other drawback is that every action requires two clicks, which can be annoying for the user.

### Execution - Undo pattern

Another way to handle this is to execute the action and allow the user to undo it. This is the pattern used by Gmail, for example.

This pattern is more complex to implement but avoids both drawbacks from the first option. The user can perform actions with only one click and the functionality is also available in API calls.

One potential drawback is that the GUI needs to understand what actions can be undone and how to undo them. This can be solved by letting the server return the action that needs to be executed to undo the previous one.


## Decision Outcome

We will use the execution - undo pattern to handle user actions that have side effects. 

Ideally the server will describe the `undo` request for a given action, but that will be decided and designed in the architecture document.