
data "datapower_httpinputconversionmap" "test" {
  depends_on = [datapower_httpinputconversionmap.test]
  app_domain = "acc_test_domain"
}