
resource "datapower_opentelemetrysampler" "test" {
  id           = "ResTestOpenTelemetrySampler"
  app_domain   = "acceptance_test"
  parent_based = true
  type         = "always-on"
}