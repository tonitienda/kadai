Feature: Add tasks

  Scenario: Add a task without a title and description should return a validation error
    Given Alice is logged in
    When Alice adds a task with title "" and description ""
    Then the request should be refused because the data is not valid
    And Alice should have 0 tasks


  Scenario: Add a task without a title should return a validation error
    Given Alice is logged in
    When Alice adds a task with title "" and description "This is the description"
    Then the request should be refused because the data is not valid
    And Alice should have 0 tasks


  Scenario: Add a task without a description should return a validation error
    Given Alice is logged in
    When Alice adds a task with title "This is the title" and description ""
    Then the request should be refused because the data is not valid

  Scenario: Add a task with a valid title and description should add the task
    Given Alice is logged in
    When Alice adds a task with title "This is the title" and description "This is the description"
    Then the request should be successful
    And Alice should have 1 task
