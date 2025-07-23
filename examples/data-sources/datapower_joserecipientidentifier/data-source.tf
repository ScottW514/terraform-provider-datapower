
data "datapower_joserecipientidentifier" "test" {
  depends_on = [datapower_joserecipientidentifier.test]
  app_domain = "acc_test_domain"
}