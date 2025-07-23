
data "datapower_webappfw" "test" {
  depends_on = [datapower_webappfw.test]
  app_domain = "acc_test_domain"
}