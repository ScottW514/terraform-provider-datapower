
data "datapower_domainavailability" "test" {
  depends_on = [datapower_domainavailability.test]
  app_domain = "acc_test_domain"
}