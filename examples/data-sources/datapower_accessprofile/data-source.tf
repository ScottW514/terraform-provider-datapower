
data "datapower_accessprofile" "test" {
  depends_on = [datapower_accessprofile.test]
  app_domain = "acc_test_domain"
}