# VirusMonitoring entity test

require "minitest/autorun"
require "json"
require_relative "../Virenmonitoring_sdk"
require_relative "runner"

class VirusMonitoringEntityTest < Minitest::Test
  def test_create_instance
    testsdk = VirenmonitoringSDK.test(nil, nil)
    ent = testsdk.VirusMonitoring(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = virus_monitoring_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "virus_monitoring." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set VIRENMONITORING_TEST_VIRUS_MONITORING_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    virus_monitoring_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.virus_monitoring")))
    virus_monitoring_ref01_data = nil
    if virus_monitoring_ref01_data_raw.length > 0
      virus_monitoring_ref01_data = Helpers.to_map(virus_monitoring_ref01_data_raw[0][1])
    end

    # LIST
    virus_monitoring_ref01_ent = client.VirusMonitoring(nil)
    virus_monitoring_ref01_match = {}

    virus_monitoring_ref01_list_result = virus_monitoring_ref01_ent.list(virus_monitoring_ref01_match, nil)
    assert virus_monitoring_ref01_list_result.is_a?(Array)

  end
end

def virus_monitoring_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "virus_monitoring", "VirusMonitoringTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = VirenmonitoringSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["virus_monitoring01", "virus_monitoring02", "virus_monitoring03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["VIRENMONITORING_TEST_VIRUS_MONITORING_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "VIRENMONITORING_TEST_VIRUS_MONITORING_ENTID" => idmap,
    "VIRENMONITORING_TEST_LIVE" => "FALSE",
    "VIRENMONITORING_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["VIRENMONITORING_TEST_VIRUS_MONITORING_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["VIRENMONITORING_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
      },
      extra || {},
    ])
    client = VirenmonitoringSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["VIRENMONITORING_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["VIRENMONITORING_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
