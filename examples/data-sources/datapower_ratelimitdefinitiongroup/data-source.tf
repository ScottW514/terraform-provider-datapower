
data "datapower_ratelimitdefinitiongroup" "test" {
  depends_on = [datapower_ratelimitdefinitiongroup.test]
  app_domain = "acc_test_domain"
}