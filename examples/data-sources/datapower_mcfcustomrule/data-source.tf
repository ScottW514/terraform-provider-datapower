
data "datapower_mcfcustomrule" "test" {
  depends_on = [datapower_mcfcustomrule.test]
  app_domain = "acc_test_domain"
}