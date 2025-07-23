
data "datapower_xacmlpdp" "test" {
  depends_on = [datapower_xacmlpdp.test]
  app_domain = "acc_test_domain"
}