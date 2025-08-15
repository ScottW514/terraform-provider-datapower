
resource "datapower_open_telemetry_sampler" "test" {
  id           = "ResTestOpenTelemetrySampler"
  app_domain   = "acceptance_test"
  parent_based = true
  type         = "always-on"
}