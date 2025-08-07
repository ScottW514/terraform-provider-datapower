
resource "datapower_cryptokerberoskeytab" "test" {
  id         = "ResTestCryptoKerberosKeytab"
  app_domain = "acceptance_test"
  filename   = "cert:///kerberos-keytab"
}