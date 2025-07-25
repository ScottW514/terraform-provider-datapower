
data "datapower_opentelemetry" "test" {
  depends_on = [datapower_opentelemetry.test]
  app_domain = "acc_test_domain"
}