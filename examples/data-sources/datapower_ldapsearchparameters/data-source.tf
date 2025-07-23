
data "datapower_ldapsearchparameters" "test" {
  depends_on = [datapower_ldapsearchparameters.test]
  app_domain = "acc_test_domain"
}