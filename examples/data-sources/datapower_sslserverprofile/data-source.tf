
data "datapower_sslserverprofile" "test" {
  depends_on = [datapower_sslserverprofile.test]
  app_domain = "acc_test_domain"
}