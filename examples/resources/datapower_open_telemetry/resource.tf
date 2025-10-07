
resource "datapower_open_telemetry" "test" {
  id         = "ResTestOpenTelemetry"
  app_domain = "acceptance_test"
  exporter   = "AccTest_OpenTelemetryExporter"
  sampler    = "AccTest_OpenTelemetrySampler"
  resource_attribute = [{
    value = "value"
  }]
}