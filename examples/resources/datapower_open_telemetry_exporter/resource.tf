
resource "datapower_open_telemetry_exporter" "test" {
  id         = "ResTestOpenTelemetryExporter"
  app_domain = "acceptance_test"
  type       = "http"
  host_name  = "localhost"
  port       = 4318
  header = [{
    header_value = "value"
  }]
  processor = "batch"
  proxy_policies = [{
    reg_exp        = "*"
    skip           = false
    remote_address = "10.10.10.10"
    remote_port    = 8888
  }]
}