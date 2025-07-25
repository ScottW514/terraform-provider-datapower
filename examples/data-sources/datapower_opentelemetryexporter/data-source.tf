
data "datapower_opentelemetryexporter" "test" {
  depends_on = [datapower_opentelemetryexporter.test]
  app_domain = "acc_test_domain"
}