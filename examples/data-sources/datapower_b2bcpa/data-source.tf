
data "datapower_b2bcpa" "test" {
  depends_on = [datapower_b2bcpa.test]
  app_domain = "acc_test_domain"
}