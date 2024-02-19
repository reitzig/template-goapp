Feature: Say hello
  Scenario: Greet some person
    When I run `template-goapp hello Stuart`
    Then the exit status should be 0
    And the output should contain "Hello, Stuart!"

  Scenario: Greet a dear friend
    When I run `template-goapp hello --dear Kevin`
    Then the exit status should be 0
    And the output should contain "Hello, dear Kevin!"
