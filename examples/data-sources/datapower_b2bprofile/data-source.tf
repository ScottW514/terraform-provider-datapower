
data "datapower_b2bprofile" "test" {
  depends_on = [datapower_b2bprofile.test]
  app_domain = "acc_test_domain"
}