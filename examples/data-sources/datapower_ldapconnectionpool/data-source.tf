
data "datapower_ldapconnectionpool" "test" {
  depends_on = [datapower_ldapconnectionpool.test]
  app_domain = "acc_test_domain"
}