
data "datapower_ratelimitdefinition" "test" {
  depends_on = [datapower_ratelimitdefinition.test]
  app_domain = "acc_test_domain"
}