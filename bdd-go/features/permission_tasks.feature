Feature: Access the list of tasks

  Scenario: Access without login should return a permission error
    Given Alice is not logged in
    When Alice requests the list of tasks
    Then the list of tasks should be empty
    And the request should be refused because the user is unauthorized



  Scenario: Access after login in should return a successful response
    Given Alice is logged in
    When Alice requests the list of tasks
    Then the list of tasks should be empty
    And the request should be successful