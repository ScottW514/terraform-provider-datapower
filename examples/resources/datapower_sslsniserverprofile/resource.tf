
resource "datapower_sslsniserverprofile" "test" {
  id                 = "ResTestSSLSNIServerProfile"
  app_domain         = "acceptance_test"
  sni_server_mapping = "AccTest_SSLSNIMapping"
}