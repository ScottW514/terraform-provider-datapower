
resource "datapower_ssl_sni_server_profile" "test" {
  id                 = "ResTestSSLSNIServerProfile"
  app_domain         = "acceptance_test"
  sni_server_mapping = "AccTest_SSLSNIMapping"
}