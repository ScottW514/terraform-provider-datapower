
resource "datapower_sslsnimapping" "test" {
  id         = "SSLSNIMapping_name"
  app_domain = "acc_test_domain"
  sni_mapping = [{
    host_name_wildmat = "hostname_wildmat"
    ssl_server        = "TestAccSSLServerProfile"
  }]
}