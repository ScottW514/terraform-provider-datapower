
resource "datapower_opentelemetryexporter" "test" {
  id         = "OpenTelemetryExporter_name"
  app_domain = "acc_test_domain"
  type       = "http"
  host_name  = "localhost"
  port       = 4318
  processor  = "batch"
}