
data "datapower_sslclientprofile" "test" {
  depends_on = [datapower_sslclientprofile.test]
  app_domain = "acc_test_domain"
}