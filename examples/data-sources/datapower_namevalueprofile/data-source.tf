
data "datapower_namevalueprofile" "test" {
  depends_on = [datapower_namevalueprofile.test]
  app_domain = "acc_test_domain"
}