
data "datapower_josesignatureidentifier" "test" {
  depends_on = [datapower_josesignatureidentifier.test]
  app_domain = "acc_test_domain"
}