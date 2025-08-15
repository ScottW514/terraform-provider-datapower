
resource "datapower_open_telemetry_exporter" "test" {
  id         = "ResTestOpenTelemetryExporter"
  app_domain = "acceptance_test"
  type       = "http"
  host_name  = "localhost"
  port       = 4318
  processor  = "batch"
}