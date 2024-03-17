Feature: Provide a reasonable CLI API
  Scenario: Print version
    When I run `template-goapp --version`
    Then the exit status should be 0
    And  the output should contain "template-goapp"
    And  the output should contain "v1.2.3"

  # TODO: What else should info do?
  #       version?
  #       compiler args?
  #       ... ?

  Scenario: Print config values and source
    When I run `template-goapp info`
    Then the output should match %r<debug\s+false\s+default>
    Then the output should match %r<foo\s+bar\s+default>

  Scenario: Read config from ~/.config/template-goapp.toml
    Given a file named "/home/test/.template-goapp.toml" with:
      """
      foo = "user first"
      """
    When I run `template-goapp info`
    Then the output should match %r<foo\s+user first\s+file /home/test/.template-goapp.toml>

  Scenario: Prefer config from .config/template-goapp.toml
    Given a file named "/home/test/.template-goapp.toml" with:
      """
      foo = "user first"
      """
    And a file named ".config/template-goapp.toml" with:
      """
      foo = "local ninja"
      """
    When I run `template-goapp info`
    Then the output should match %r<foo\s+local ninja\s+file .config/template-goapp.toml>

  Scenario: Prefer config from .template-goapp.toml
    Given a file named "/home/test/.template-goapp.toml" with:
      """
      foo = "user first"
      """
    And a file named ".config/template-goapp.toml" with:
      """
      foo = "local ninja"
      """
    And a file named ".template-goapp.toml" with:
      """
      foo = "local hero"
      """
    When I run `template-goapp info`
    Then the output should match /foo\s+local hero\s+file .template-goapp.toml/

  Scenario: Prefer config from ENV var
    Given a file named "/home/test/.template-goapp.toml" with:
      """
      foo = "user first"
      """
    And a file named ".config/template-goapp.toml" with:
      """
      foo = "local ninja"
      """
    And a file named ".template-goapp.toml" with:
      """
      foo = "local hero"
      """
    And I set the environment variable "TEMPLATE_GOAPP_FOO" to "warming"
    When I run `template-goapp info`
    Then the output should match /foo\s+warming\s+env var TEMPLATE_GOAPP_FOO/

  Scenario: Prefer config from CLI option
    Given a file named "/home/test/.template-goapp.toml" with:
      """
      foo = "user first"
      """
    And a file named ".config/template-goapp.toml" with:
      """
      foo = "local ninja"
      """
    And a file named ".template-goapp.toml" with:
      """
      foo = "local hero"
      """
    And I set the environment variable "TEMPLATE_GOAPP_FOO" to "warming"
    When I run `template-goapp --foo turtle info`
    Then the output should match /foo\s+turtle\s+CLI option --foo/
