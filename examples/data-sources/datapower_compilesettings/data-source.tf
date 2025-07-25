
data "datapower_compilesettings" "test" {
  depends_on = [datapower_compilesettings.test]
  app_domain = "acc_test_domain"
}