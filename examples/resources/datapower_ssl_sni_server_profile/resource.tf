
resource "datapower_ssl_sni_server_profile" "test" {
  id         = "ResTestSSLSNIServerProfile"
  app_domain = "acceptance_test"
  protocols = {
    ssl_v3   = false
    tls_v1d0 = false
    tls_v1d1 = true
    tls_v1d2 = true
    tls_v1d3 = true
  }
  sni_server_mapping = "AccTest_SSLSNIMapping"
  ssl_options = {
  }
}