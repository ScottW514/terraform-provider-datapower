
data "datapower_mcfhttpmethod" "test" {
  depends_on = [datapower_mcfhttpmethod.test]
  app_domain = "acc_test_domain"
}