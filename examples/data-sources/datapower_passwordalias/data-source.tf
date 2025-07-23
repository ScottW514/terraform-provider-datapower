
data "datapower_passwordalias" "test" {
  depends_on = [datapower_passwordalias.test]
  app_domain = "acc_test_domain"
}