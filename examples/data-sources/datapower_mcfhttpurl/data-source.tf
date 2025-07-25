
data "datapower_mcfhttpurl" "test" {
  depends_on = [datapower_mcfhttpurl.test]
  app_domain = "acc_test_domain"
}