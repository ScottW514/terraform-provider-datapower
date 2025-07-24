
data "datapower_b2bprofilegroup" "test" {
  depends_on = [datapower_b2bprofilegroup.test]
  app_domain = "acc_test_domain"
}