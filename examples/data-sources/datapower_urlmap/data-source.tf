
data "datapower_urlmap" "test" {
  depends_on = [datapower_urlmap.test]
  app_domain = "acc_test_domain"
}