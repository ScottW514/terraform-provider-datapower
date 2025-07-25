
data "datapower_opentelemetrysampler" "test" {
  depends_on = [datapower_opentelemetrysampler.test]
  app_domain = "acc_test_domain"
}