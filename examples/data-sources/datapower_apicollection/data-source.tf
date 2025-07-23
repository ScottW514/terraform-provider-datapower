
data "datapower_apicollection" "test" {
  depends_on = [datapower_apicollection.test]
  app_domain = "acc_test_domain"
}