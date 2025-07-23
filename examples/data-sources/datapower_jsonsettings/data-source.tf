
data "datapower_jsonsettings" "test" {
  depends_on = [datapower_jsonsettings.test]
  app_domain = "acc_test_domain"
}