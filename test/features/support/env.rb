# frozen_string_literal: true

require 'aruba/cucumber'

# The module referenced by self from within scenarios
module RunnerWorld
  attr_accessor :last_command_status, :last_command_output

  def verbose?
    ENV.fetch('VERBOSE_TEST_OUTPUT', false) == 'true'
  end
end

World(RunnerWorld)

BeforeAll do
  puts "INFO: VERBOSE_TEST_OUTPUT=#{ENV.fetch('VERBOSE_TEST_OUTPUT', false)}"
end
