
resource "datapower_opentelemetry" "test" {
  id         = "OpenTelemetry_name"
  app_domain = "acc_test_domain"
  exporter   = datapower_opentelemetryexporter.test.id
  sampler    = datapower_opentelemetrysampler.test.id
}