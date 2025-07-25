
data "datapower_soapheaderdisposition" "test" {
  depends_on = [datapower_soapheaderdisposition.test]
  app_domain = "acc_test_domain"
}