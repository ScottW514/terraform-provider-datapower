
data "datapower_parsesettings" "test" {
  depends_on = [datapower_parsesettings.test]
  app_domain = "acc_test_domain"
}