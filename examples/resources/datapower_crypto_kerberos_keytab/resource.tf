
resource "datapower_crypto_kerberos_keytab" "test" {
  id         = "ResTestCryptoKerberosKeytab"
  app_domain = "acceptance_test"
  filename   = "cert:///kerberos-keytab"
  gss_checksum_flags = {
  }
}