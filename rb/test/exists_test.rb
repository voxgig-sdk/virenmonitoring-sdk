# Virenmonitoring SDK exists test

require "minitest/autorun"
require_relative "../Virenmonitoring_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = VirenmonitoringSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
