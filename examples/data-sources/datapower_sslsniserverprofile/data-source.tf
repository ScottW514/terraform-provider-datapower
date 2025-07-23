
data "datapower_sslsniserverprofile" "test" {
  depends_on = [datapower_sslsniserverprofile.test]
  app_domain = "acc_test_domain"
}