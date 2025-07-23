
data "datapower_apirouting" "test" {
  depends_on = [datapower_apirouting.test]
  app_domain = "acc_test_domain"
}