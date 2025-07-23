
data "datapower_httpuseragent" "test" {
  depends_on = [datapower_httpuseragent.test]
  app_domain = "acc_test_domain"
}