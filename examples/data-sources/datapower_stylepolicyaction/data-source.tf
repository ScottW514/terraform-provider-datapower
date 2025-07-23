
data "datapower_stylepolicyaction" "test" {
  depends_on = [datapower_stylepolicyaction.test]
  app_domain = "acc_test_domain"
}