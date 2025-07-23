
data "datapower_jwerecipient" "test" {
  depends_on = [datapower_jwerecipient.test]
  app_domain = "acc_test_domain"
}