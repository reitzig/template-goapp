# frozen_string_literal: true

require 'aruba/cucumber'

# The module referenced by self from within scenarios
module RunnerWorld
  def verbose?
    ENV.fetch('VERBOSE_TEST_OUTPUT', false) == 'true'
  end
end

World(RunnerWorld)

BeforeAll do
  if ENV.fetch('RUNNING_IN_CONTAINER', false) != 'yessir'
    Kernel.abort("We don't seem to be running in a container; aborting for safety!")
  end

  puts "INFO: VERBOSE_TEST_OUTPUT=#{ENV.fetch('VERBOSE_TEST_OUTPUT', false)}"
end
