
resource "datapower_sslsnimapping" "test" {
  id         = "ResTestSSLSNIMapping"
  app_domain = "acceptance_test"
  sni_mapping = [{
    host_name_wildmat = "hostname_wildmat"
    ssl_server        = "AccTest_SSLServerProfile"
  }]
}