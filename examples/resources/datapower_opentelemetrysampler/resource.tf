
resource "datapower_opentelemetrysampler" "test" {
  id           = "OpenTelemetrySampler_name"
  app_domain   = "acc_test_domain"
  parent_based = true
  type         = "always-on"
}