
data "datapower_mcfhttpheader" "test" {
  depends_on = [datapower_mcfhttpheader.test]
  app_domain = "acc_test_domain"
}