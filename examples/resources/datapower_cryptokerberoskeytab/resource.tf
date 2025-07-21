
resource "datapower_cryptokerberoskeytab" "test" {
  id         = "CryptoKerberosKeytab_test"
  app_domain = "acc_test_domain"
  filename   = "cert:///kerberos-keytab"
}