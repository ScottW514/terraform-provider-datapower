
data "datapower_apildapregistry" "test" {
  depends_on = [datapower_apildapregistry.test]
  app_domain = "acc_test_domain"
}