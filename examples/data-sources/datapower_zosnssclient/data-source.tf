
data "datapower_zosnssclient" "test" {
  depends_on = [datapower_zosnssclient.test]
  app_domain = "acc_test_domain"
}