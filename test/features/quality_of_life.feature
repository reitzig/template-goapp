Feature: Provide a reasonable CLI API
  Scenario: Print version
    When I run `template-goapp --version`
    Then the exit status should be 0
    And  the output should contain "template-goapp"
    And  the output should contain "v1.2.3"
