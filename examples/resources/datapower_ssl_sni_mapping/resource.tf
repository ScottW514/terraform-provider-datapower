
resource "datapower_ssl_sni_mapping" "test" {
  id         = "ResTestSSLSNIMapping"
  app_domain = "acceptance_test"
  sni_mapping = [{
    host_name_wildmat = "hostname_wildmat"
    ssl_server        = "AccTest_SSLServerProfile"
  }]
}