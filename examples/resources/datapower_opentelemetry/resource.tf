
resource "datapower_opentelemetry" "test" {
  id         = "ResTestOpenTelemetry"
  app_domain = "acceptance_test"
  exporter   = "AccTest_OpenTelemetryExporter"
  sampler    = "AccTest_OpenTelemetrySampler"
}