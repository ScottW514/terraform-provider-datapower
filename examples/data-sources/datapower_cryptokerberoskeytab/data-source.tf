
data "datapower_cryptokerberoskeytab" "test" {
  depends_on = [datapower_cryptokerberoskeytab.test]
  app_domain = "acc_test_domain"
}