
data "datapower_jwssignature" "test" {
  depends_on = [datapower_jwssignature.test]
  app_domain = "acc_test_domain"
}