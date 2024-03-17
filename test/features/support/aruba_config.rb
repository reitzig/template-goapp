# frozen_string_literal: true

Aruba.configure do |config|
  # NB: We only run these tests in a container,
  #     so this is fine. 🔥
  config.allow_absolute_paths = true
end
