
data "datapower_webspherejmsserver" "test" {
  depends_on = [datapower_webspherejmsserver.test]
  app_domain = "acc_test_domain"
}