
resource "datapower_opentelemetry" "test" {
  id         = "OpenTelemetry_name"
  app_domain = "acc_test_domain"
  exporter   = "TestAccOpenTelemetryExporter"
  sampler    = "TestAccOpenTelemetrySampler"
}